//go:build mage

package main

import (
	"fmt"

	"github.com/mcandre/buttery"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/mcandre/mx"
)

// Default references the default build task.
var Default = Test

// Audit runs security checks.
func Audit() error { return Govulncheck() }

// Clean removes artifacts.
func Clean() error { mg.Deps(CleanPackages); return CleanArtifacts() }

// CleanArtifacts removes artifacts.
func CleanArtifacts() error { return mx.CleanGoReleaser() }

// CleanPackages removes OS package artifacts.
func CleanPackages() error { return sh.RunV("rockhopper", "-c") }

// Deadcode runs deadcode.
func Deadcode() error { return sh.RunV("deadcode", "./...") }

// DockerBuild creates local Docker buildx images.
func DockerBuild() error { return sh.RunV("docker", "buildx", "bake", "all") }

// DockerPush creates and tag aliases remote Docker buildx images.
func DockerPush() error { return sh.RunV("docker", "buildx", "bake", "test", "--push") }

// DockerTest creates and tag aliases remote test Docker buildx images.
func DockerTest() error { return sh.RunV("docker", "buildx", "bake", "production", "--push") }

// Errcheck runs errcheck.
func Errcheck() error { return sh.RunV("errcheck", "-blank") }

// GoImports runs goimports.
func GoImports() error { return mx.GoImports("-w") }

// Goreleaser builds crossplatform binaries and tarballs.
func Goreleaser() error {
	return mx.GoReleaser(
		map[string]string{"NO_COLOR": "1"},
		"--snapshot", "--clean",
	)
}

// GoVet runs default go vet analyzers.
func GoVet() error { return mx.GoVet() }

// Govulncheck runs govulncheck.
func Govulncheck() error { return sh.RunV("govulncheck", "-scan", "package", "./...") }

// Install builds and installs Go applications.
func Install() error { return mx.Install() }

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

// Uninstall deletes installed Go applications.
func Uninstall() error { return mx.Uninstall("buttery") }

// Upload sends packages to CloudFlare R2.
func Upload() error { return sh.RunV("./upload") }
