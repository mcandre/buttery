# BUILDTIME REQUIREMENTS

* a UNIX-like environment (e.g. [WSL](https://learn.microsoft.com/en-us/windows/wsl/))
* [Go](https://go.dev/)
* POSIX compliant [make](https://pubs.opengroup.org/onlinepubs/9799919799/utilities/make.html)
* [Rust](https://rust-lang.org/)
* Provision additional dev tools with `make`

## Recommended

* [ASDF](https://asdf-vm.com/) 0.18 (run `asdf reshim` after provisioning)
* [direnv](https://direnv.net/) 2
* [ffmpeg](https://ffmpeg.org/) edits and converts videos
* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's
* [mkvtools](https://emmgunn.com/wp/mkvtools-home/) edits MKV videos
* macOS [open](https://ss64.com/mac/open.html) or equivalent alias
* [webm](https://www.webmproject.org/) supports audio in animation loops

## Windows

Apply a user environment variable `GODEBUG=modcacheunzipinplace=1` per [access denied resolution](https://github.com/golang/go/wiki/Modules/e93463d3e853031af84204dc5d3e2a9a710a7607#go-115), for native Windows development environments (Command Prompt / PowerShell, not WLS, not Cygwin, not MSYS2, not MinGW, not msysGit, not Git Bash, not etc).

# GENERATE SOURCES

After each change to an enum, regenerate auxiliary Go sources:

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

# PORT

```sh
mage port
```

# TEST DOCKER IMAGES

```sh
mage dockerTest
```

# PUSH DOCKER IMAGES

```sh
mage dockerPush
```
