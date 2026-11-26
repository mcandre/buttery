# NATIVE DEPENDENCIES

Some Rust dependencies rely on native C/C++ resources.

## Tools

* [autotools](https://www.gnu.org/software/automake/manual/html_node/Autotools-Introduction.html)
* [clang](https://clang.llvm.org/)

## C/C++ Headers

* [BLAKE2](https://www.blake2.net/)
* [bzip2](https://sourceware.org/bzip2/)

## Debian

```sh
sudo apt-get update
sudo apt-get install -y \
    build-essential \
    clang \
    libacl1-dev \
    libb2-dev \
    libbz2-dev \
    liblz4-dev \
    libzstd-dev
```

## FreeBSD

```sh
sudo pkg install -y \
    autoconf \
    automake \
    libb2 \
    libtool \
    liblz4 \
    gmake
```

## macOS (Homebrew)

```sh
xcode-select --install
brew install -y \
    bzip2 \
    libb2 \
    lz4 \
    zstd
```

## RHEL

```sh
sudo dnf install -y \
    @development-tools \
    clang \
    libacl1-devel \
    libb2-devel \
    libbz2-devel \
    libzstd-devel \
    lz4-devel
```

# FAQ

## `error: libarchive error (code -1): Using external zstd program [errno: -1]`

### Problem

libarchive2-sys may attempt to shell out to zstd, such as when building Debian family packages.

This can happen when the zstd (header) package is missing. Or was missing at the time when chiri's dependencies were initially compiled.

### Solution

Fully rebuild chiri, including dependencies, in a way that targets the zstd library, rather than zstd executables.

1. Ensure that your operating system's zstd header package is installed.
2. Run `cargo clean` to reset cargo's wrecked libarchive2-sys cache.
3. Recompile chiri.
