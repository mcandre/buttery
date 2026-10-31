package gifsicle

import (
	"bytes"
	"embed"
	"fmt"
	"image/gif"
	"io"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/munchpass/gifsicle-go/embedbinwrapper"
)

//go:embed bin/*
var binariesFs embed.FS

func createBinWrapper(binaryName string) (*embedbinwrapper.EmbedBinWrapper, error) {
	b := embedbinwrapper.NewExecutableBinWrapper()

	switch runtime.GOOS {
	case "windows":
		binPath := fmt.Sprintf("bin/win/x64/%s", binaryName)
		ext := strings.ToLower(filepath.Ext(binPath))
		if ext != ".exe" {
			binPath += ".exe"
		}

		binary, err := binariesFs.ReadFile(binPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read embed binary: %s", err)
		}
		return b.Src(embedbinwrapper.NewSrc().Bin(binary).Os("win32")), nil
	case "linux":
		binary, err := binariesFs.ReadFile(fmt.Sprintf("bin/linux/x64/%s", binaryName))
		if err != nil {
			return nil, fmt.Errorf("failed to read embed binary: %s", err)
		}
		return b.Src(embedbinwrapper.NewSrc().Bin(binary).Os("linux")), nil
	case "darwin":
		binary, err := binariesFs.ReadFile(fmt.Sprintf("bin/macos/%s", binaryName))
		if err != nil {
			return nil, fmt.Errorf("failed to read embed binary: %s", err)
		}
		return b.Src(embedbinwrapper.NewSrc().Bin(binary).Os("darwin")), nil
	default:
		return nil, fmt.Errorf("unsupported OS %s", runtime.GOOS)
	}
}

func createReaderFromGif(g *gif.GIF) (io.Reader, error) {
	var buffer bytes.Buffer
	err := gif.EncodeAll(&buffer, g)
	return &buffer, err
}

func version(b *embedbinwrapper.EmbedBinWrapper) (string, error) {
	b.Reset()
	err := b.Run("--version")

	if err != nil {
		return "", err
	}

	v := string(b.StdOut())
	return v, nil
}
