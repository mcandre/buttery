# INSTALL GUIDE

In addition to OS packages, buttery also supports alternative installation methods.

# INSTALL (CURL)

curl based installs automatically download and extract precompiled binaries.

```sh
curl -L https://raw.githubusercontent.com/mcandre/buttery/refs/heads/main/install-buttery | sh
```

## Postinstall

Ensure `$HOME/.local/bin` is registered with your shell's `PATH` environment variable.

## Uninstall

```sh
curl -L https://raw.githubusercontent.com/mcandre/buttery/refs/heads/main/uninstall-buttery | sh
```

## System Requirements

### Bitness

64

### Operating Systems

* DragonflyBSD (Intel)
* FreeBSD (ARM, Intel)
* Illumos (Intel)
* Linux (ARM, Intel, LoongArch, RISC-V)
* macOS 26 Tahoe+ (ARM, Intel)
* NetBSD (ARM, Intel)
* OpenBSD (ARM, Intel, RISC-V)
* WSL 2 (ARM, Intel)

### Prerequisites

* [bash](https://www.gnu.org/software/bash/) 4+
* [curl](https://curl.se/)

# INSTALL (PRECOMPILED BINARIES)

Precompiled binaries may be installed manually.

## Install

1. Download a [tarball](https://github.com/mcandre/buttery/releases) corresponding to your environment's architecture and OS.
2. Extract executables into a selected directory.

   Examples:

   * `~/.local/bin` (XDG compliant per-user)
   * `/usr/local/bin` (XDG compliant global)
   * `~/bin` (BSD)
   * `~\AppData\Local` (native Windows)

## Postinstall

Ensure the selected directory is registered with your shell's `PATH` environment variable.

## Uninstall

Remove the application executables from the selected directory.

## System Requirements

### Bitness

64

### Operating Systems

* DragonflyBSD (Intel)
* FreeBSD (ARM, Intel)
* Illumos (Intel)
* Linux (ARM, Intel, LoongArch, RISC-V)
* macOS 26 Tahoe+ (ARM, Intel)
* NetBSD (ARM, Intel)
* OpenBSD (ARM, Intel, RISC-V)
* Windows 11+ (ARM, Intel)

# INSTALL (GO REMOTE)

buttery is packaged as a Go module.

```sh
go install github.com/mcandre/buttery/cmd/buttery@latest
```

## Prerequisites

* [Go](https://go.dev/)

# INSTALL (GO LOCAL)

buttery may be compiled from source.

```sh
git clone https://github.com/mcandre/buttery.git
cd buttery
go install ./...
```

## Prerequisites

* [git](https://git-scm.com/)
* [Go](https://go.dev/)

For more details on developing buttery, see our [development guide](DEVELOPMENT.md).
