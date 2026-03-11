//go:build mage

package main

import (
	"github.com/mcandre/buttery"
	"github.com/magefile/mage/mg"
	mageextras "github.com/mcandre/mage-extras"

	"fmt"
)

// Default references the default build task.
var Default = Test

// Audit runs security checks.
func Audit() error { return Govulncheck() }

// Clean removes artifacts.
func Clean() error { mg.Deps(CleanPackages); return CleanArtifacts() }

// CleanArtifacts removes artifacts.
func CleanArtifacts() error { return mageextras.Run("tuco", "-clean") }

// CleanPackages removes OS package artifacts.
func CleanPackages() error { return mageextras.Run("rockhopper", "-c") }

// Deadcode runs deadcode.
func Deadcode() error { return mageextras.Run("deadcode", "./...") }

// DockerBuild creates local Docker buildx images.
func DockerBuild() error { return mageextras.Run("docker", "buildx", "bake", "all") }

// DockerPush creates and tag aliases remote Docker buildx images.
func DockerPush() error { return mageextras.Run("docker", "buildx", "bake", "test", "--push") }

// DockerTest creates and tag aliases remote test Docker buildx images.
func DockerTest() error { return mageextras.Run("docker", "buildx", "bake", "production", "--push") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Run("errcheck", "-blank") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// GoVet runs default go vet analyzers.
func GoVet() error { return mageextras.GoVet() }

// Govulncheck runs govulncheck.
func Govulncheck() error { return mageextras.Run("govulncheck", "-scan", "package", "./...") }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

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
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Package generates OS packages.
func Package() error { return mageextras.Run("rockhopper", "-r", fmt.Sprintf("version=%s", buttery.Version)) }

// Shadow runs go vet with shadow checks enabled.
func Shadow() error { return mageextras.GoVetShadow() }

// Staticcheck runs staticcheck.
func Staticcheck() error { return mageextras.Run("staticcheck", "./...") }

// Test executes a test suite.
func Test() error { return mageextras.UnitTest() }

// Tuco builds crossplatform binaries and tarballs.
func Tuco() error { return mageextras.Run("tuco") }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("buttery") }
