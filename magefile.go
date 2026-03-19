//go:build mage

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/mcandre/buttery"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mcandre/mx"
)

// ArtifactsPath describes where artifacts are produced.
const ArtifactsPath = "bin"

// Default references the default build task.
var Default = Build

// Audit runs security checks.
func Audit() error { return Govulncheck() }

// Build compiles Go projects.
func Build() error {
	dest := ArtifactsPath

	if d, ok := os.LookupEnv("DEST"); ok && d != "" {
		dest = d
	}

	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	return sh.RunV("go", "build", "-o", dest, "./...")
}

// Clean removes artifacts.
func Clean() error { mg.Deps(CleanPackages); mg.Deps(CleanBuild); return CleanArtifacts() }

// CleanArtifacts removes artifacts.
func CleanArtifacts() error { return sh.RunV("tuco", "-clean") }

// CleanBuild removes build artifacts.
func CleanBuild() error { return os.RemoveAll(ArtifactsPath) }

// CleanPackages removes OS package artifacts.
func CleanPackages() error { return sh.RunV("rockhopper", "-c") }

// Deadcode runs deadcode.
func Deadcode() error { return sh.RunV("deadcode", "./...") }

// Errcheck runs errcheck.
func Errcheck() error { return sh.RunV("errcheck", "-blank") }

// GoGenerate populates generated Go source code.
func GoGenerate() error { return sh.RunV("go", "generate", "./...") }

// GoImports runs goimports.
func GoImports() error { return mx.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mx.GoVet() }

// Govulncheck runs govulncheck.
func Govulncheck() error { return sh.RunV("govulncheck", "-scan", "package", "./...") }

// Install builds and installs Go applications.
func Install() error { mg.Deps(GoGenerate); return mx.Install() }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(Deadcode)
	mg.Deps(GoImports)
	mg.Deps(GoVet)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	mg.Deps(Shadow)
	mg.Deps(Staticcheck)
	return nil
}

// Nakedret runs nakedret.
func Nakedret() error { return mx.Nakedret("-l", "0") }

// Package generates OS packages.
func Package() error { return sh.RunV("rockhopper", "-r", fmt.Sprintf("version=%s", buttery.Version)) }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mx.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return sh.RunV("staticcheck", "./...") }

// Test executes a test suite.
func Test() error { return mx.UnitTest() }

// Tuco builds crossplatform binaries and tarballs.
func Tuco() error { return sh.RunV("tuco") }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mx.Uninstall("buttery") }

// Bucket stores OS packages
const Bucket = "s3://buttery"

// Artifacts contains precompiled binaries
var Artifacts = path.Join(".rockhopper", "artifacts")

// Banner identifies the application version.
var Banner = fmt.Sprintf("buttery-%s", buttery.Version)

// Dest stores OS packages for this application version.
var Dest = fmt.Sprintf("%s/%s/", Bucket, Banner)

// Upload sends packages to CloudFlare R2.
func Upload() error {
	return mx.RunVSilent("aws",
		"--cli-connect-timeout", "1",
		"s3",
		"cp",
		"--recursive",
		Artifacts,
		Dest,
	)
}
