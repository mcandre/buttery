# DEVELOPMENT GUIDE

buttery follows standard, Go based operations for compiling and unit testing Go code.

For advanced operations, such as linting, we further supplement with some software industry tools.

# BUILDTIME REQUIREMENTS

## Prerequisites

* a UNIX-like environment (e.g. [WSL](https://learn.microsoft.com/en-us/windows/wsl/))
* [awscli](https://aws.amazon.com/cli/)
* [bash](https://www.gnu.org/software/bash/) 4+
* [Docker](https://www.docker.com/)
* [Go](https://go.dev/)
* POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* [Rust](https://rust-lang.org/)
* Provision additional dev tools with `make`

## Recommended

* [asdf](https://asdf-vm.com/) 0.18
* macOS [open](https://ss64.com/mac/open.html) or equivalent alias

## Postinstall

Register output of `go env GOBIN` to `PATH` environment variable.

# GENERATE SOURCES

After each change to `stitch.go`, regenerate auxiliary Go sources:

```sh
stringer -type "Stitch"
```

# AUDIT

```sh
mage audit
```

# INSTALL

```sh
mage install
```

# UNINSTALL

```sh
mage uninstall
```

# LINT

```sh
mage lint
```

# TEST

```sh
mage test
```

# CROSSCOMPILE & ARCHIVE BINARIES

```sh
mage tuco
```

# PACKAGE BINARIES

```sh
mage package
```

# UPLOAD OS PACKAGES

```sh
mage upload
```

# CLEAN

```sh
mage clean
```
