# BUILDTIME REQUIREMENTS

* [Go](https://go.dev/) 1.23.2+
* POSIX compatible [make](https://pubs.opengroup.org/onlinepubs/9699919799/utilities/make.html)
* [Rust](https://www.rust-lang.org/) 1.75.0+
* [Snyk](https://snyk.io/)
* POSIX compatible [tar](https://pubs.opengroup.org/onlinepubs/7908799/xcu/tar.html)
* Provision additional dev tools with `make`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.10 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2
* [ffmpeg](https://ffmpeg.org/) edits and converts videos
* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's
* [mkvtools](https://emmgunn.com/wp/mkvtools-home/) edits MKV videos
* macOS [open](https://ss64.com/mac/open.html) or equivalent alias
* [webm](https://www.webmproject.org/) supports audio in animation loops
* a UNIX environment, such as macOS, Linux, BSD, [WSL](https://learn.microsoft.com/en-us/windows/wsl/), etc.

Non-UNIX environments may produce subtle adverse effects when linting or generating application ports.

## Windows

Apply a user environment variable `GODEBUG=modcacheunzipinplace=1` per [access denied resolution](https://github.com/golang/go/wiki/Modules/e93463d3e853031af84204dc5d3e2a9a710a7607#go-115), for native Windows development environments (Command Prompt / PowerShell, not WLS, not Cygwin, not MSYS2, not MinGW, not msysGit, not Git Bash, not etc).

# GENERATE SOURCES

After each change to an enum, regenerate auxiliary Go sources:

```console
$ stringer ./...
```

# AUDIT

```console
$ mage audit
```

# INSTALL

```console
$ mage install
```

# UNINSTALL

```console
$ mage uninstall
```

# LINT

```console
$ mage lint
```

# TEST

```console
$ mage test
```

# PORT

```console
$ mage port
```
