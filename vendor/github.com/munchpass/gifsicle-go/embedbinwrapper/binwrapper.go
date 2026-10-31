package embedbinwrapper

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/amenzhinsky/go-memexec"
)

/*
Wraps an executable binary and runs it in memory.

Inspired by https://github.com/nickalie/go-binwrapper.
*/
type EmbedBinWrapper struct {
	allSrc []*Src

	// TODO: why is this not used?
	// output   []byte

	stdErr       []byte
	stdOut       []byte
	stdIn        io.Reader
	stdOutWriter io.Writer

	args    []string
	env     []string
	debug   bool
	cmd     *exec.Cmd
	timeout time.Duration
}

// NewExecutableBinWrapper creates ExecutableBinWrapper instance
func NewExecutableBinWrapper() *EmbedBinWrapper {
	return &EmbedBinWrapper{}
}

// Src adds a Src to BinWrapper
func (b *EmbedBinWrapper) Src(src *Src) *EmbedBinWrapper {
	b.allSrc = append(b.allSrc, src)
	return b
}

// Timeout sets timeout for the command. By default it's 0 (binary will run till end).
func (b *EmbedBinWrapper) Timeout(timeout time.Duration) *EmbedBinWrapper {
	b.timeout = timeout
	return b
}

/*
Arg adds command line argument to run the binary with.

Specify values only if you have an argument that is separated by a space.

For instance, if you want `-o xyz`, do:

	b.Arg("-o", "xyz")

Otherwise, if you want `--output=xyz`, do:

	b.Arg("--output=xyz")
*/
func (b *EmbedBinWrapper) Arg(name string, values ...string) *EmbedBinWrapper {
	values = append([]string{name}, values...)
	b.args = append(b.args, values...)
	return b
}

// Debug enables debug output
func (b *EmbedBinWrapper) Debug() *EmbedBinWrapper {
	b.debug = true
	return b
}

// Args returns arguments were added with Arg method
func (b *EmbedBinWrapper) Args() []string {
	return b.args
}

// StdIn sets reader to read executable's stdin from
func (b *EmbedBinWrapper) StdIn(reader io.Reader) *EmbedBinWrapper {
	b.stdIn = reader
	return b
}

// StdOut returns the binary's stdout after Run was called
func (b *EmbedBinWrapper) StdOut() []byte {
	return b.stdOut
}

// CombinedOutput returns combined executable's stdout and stderr
func (b *EmbedBinWrapper) CombinedOutput() []byte {
	return append(b.stdOut, b.stdErr...)
}

// SetStdOut set writer to write executable's stdout
func (b *EmbedBinWrapper) SetStdOut(writer io.Writer) *EmbedBinWrapper {
	b.stdOutWriter = writer
	return b
}

// Env specifies the environment of the executable.
// If Env is nil, Run uses the current process's environment.
// Elements of env should be in form: "ENV_VARIABLE_NAME=value"
func (b *EmbedBinWrapper) Env(env []string) *EmbedBinWrapper {
	b.env = env
	return b
}

// StdErr returns the executable's stderr after Run was called
func (b *EmbedBinWrapper) StdErr() []byte {
	return b.stdErr
}

// Reset removes all arguments set with Arg method, cleans StdOut and StdErr
func (b *EmbedBinWrapper) Reset() *EmbedBinWrapper {
	b.args = []string{}
	b.stdOut = nil
	b.stdErr = nil
	b.stdIn = nil
	b.stdOutWriter = nil
	b.env = nil
	b.cmd = nil
	return b
}

func stringsContains(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}

func osFilterObj(values []*Src) *Src {
	arches := []string{runtime.GOARCH}

	if runtime.GOARCH == "386" {
		arches = append(arches, "x86")
	} else if runtime.GOARCH == "amd64" {
		arches = append(arches, "x64")
	}

	platforms := []string{runtime.GOOS}

	if runtime.GOOS == "windows" {
		platforms = append(platforms, "win32")
	}

	for _, v := range values {
		if stringsContains(platforms, v.os) && stringsContains(arches, v.arch) {
			return v
		} else if stringsContains(platforms, v.os) && v.arch == "" {
			return v
		} else if stringsContains(arches, v.arch) && v.os == "" {
			return v
		} else if v.os == "" && v.arch == "" {
			return v
		}
	}

	return nil
}

// Goes through all of the binary sources and matches
// with the source that most closely matches the current OS and Architecture
func (b *EmbedBinWrapper) findMatchingBinarySrc() (*Src, error) {
	src := osFilterObj(b.allSrc)
	if src == nil {
		return nil, errors.New("no binary source found matching your OS/architecture. It's probably not supported")
	}

	return src, nil
}

// Run runs the binary with provided arg list.
// Arg list is appended to args set through Arg method
// Returns context.DeadlineExceeded in case of timeout
func (b *EmbedBinWrapper) Run(arg ...string) error {
	if len(b.allSrc) == 0 {
		return fmt.Errorf("need at least one binary source to run")
	}

	matchedSrc, err := b.findMatchingBinarySrc()
	if err != nil {
		return err
	}

	binExecutable, err := memexec.New(matchedSrc.bin)
	if err != nil {
		return err
	}

	defer binExecutable.Close()

	arg = append(b.args, arg...)

	if b.debug {
		fmt.Println("BinWrapper.Run: [executable_path] " + strings.Join(arg, " "))
	}

	var ctx context.Context
	var cancel context.CancelFunc
	if b.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), b.timeout)
	} else {
		ctx = context.Background()
		cancel = func() {}
	}
	defer cancel()

	b.cmd = binExecutable.CommandContext(ctx, arg...)

	if b.env != nil {
		b.cmd.Env = b.env
	}

	if b.stdIn != nil {
		b.cmd.Stdin = b.stdIn
	}

	var stdout io.Reader

	if b.stdOutWriter != nil {
		b.cmd.Stdout = b.stdOutWriter
	} else {
		stdout, _ = b.cmd.StdoutPipe()
	}

	stderr, _ := b.cmd.StderrPipe()

	err = b.cmd.Start()

	if err != nil {
		return err
	}

	if stdout != nil {
		b.stdOut, _ = io.ReadAll(stdout)
	}

	b.stdErr, _ = io.ReadAll(stderr)
	err = b.cmd.Wait()

	if ctx.Err() == context.DeadlineExceeded {
		return context.DeadlineExceeded
	}

	return err
}

// Kill terminates the process
func (b *EmbedBinWrapper) Kill() error {
	if b.cmd != nil && b.cmd.Process != nil {
		return b.cmd.Process.Kill()
	}

	return nil
}
