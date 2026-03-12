package mx

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/magefile/mage/sh"
)

// CleanGoReleaser removes artifacts from the standard goreleaser directory "dist".
func CleanGoReleaser() error { return sh.Rm("dist") }

// GoReleaser executes goreleaser with the UNIX / Go idiom of silencing extraneous output.
func GoReleaser(env map[string]string, args ...string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	_, err := sh.Exec(env, &stdout, &stderr, "goreleaser", args...)

	if err != nil {
		if _, err2 := io.Copy(os.Stdout, &stdout); err2 != nil {
			log.Println(err2)
		}

		if _, err2 := io.Copy(os.Stderr, &stderr); err2 != nil {
			log.Println(err2)
		}
	}

	return err
}
