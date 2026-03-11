# mage-extras: some predefined tasks for common mage workflows

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/mage-extras) [![Test](https://github.com/mcandre/mage-extras/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/mage-extras/actions/workflows/test.yml) [![license](https://img.shields.io/badge/license-BSD-0)](LICENSE.md)

# SUMMARY

mage-extras streamlines common Go development tasks.

# ABOUT

[API Docs](https://pkg.go.dev/github.com/mcandre/mage-extras)

mage-extras provides utility functions for common Go development operations.

* `Install` - Compile and install Go executables
* Lint Go projects recursively:
  * `GoImports`
  * `GoLint`
  * `GoVet`
  * `GoVetShadow`
  * `Nakedret`
  * `Revive`
* `UnitTest` - trigger unit test suite
* `Xgo` crosscompile cGo applications
* `Run` for everything else

# SYSTEM REQUIREMENTS

* [Go](https://go.dev/)
* [Mage](https://magefile.org/) (e.g., `go get -tool github.com/magefile/mage`)

For details on developing mage-extras, see our [development guide](DEVELOPMENT.md).
