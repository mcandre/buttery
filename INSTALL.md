# INSTALL GUIDE

In addition to OS packages, buttery also supports alternative installation methods.

# INSTALL (GO REMOTE)

buttery is packaged as a Go module.

```sh
go install github.com/mcandre/buttery/cmd/buttery@latest
```

## Prerequisites

* [Go](https://go.dev/)

## Postinstall

Register `"$(go env GOBIN)"` to `PATH` environment variable.

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

## Postinstall

Register `"$(go env GOBIN)"` to `PATH` environment variable.

For more details on developing buttery, see our [development guide](DEVELOPMENT.md).
