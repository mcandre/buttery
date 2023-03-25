# buttery: a video editor with motion smoothing

![examples/homer.buttery.gif](examples/homer.buttery.gif)

# ABOUT

`buttery` generates continuous GIF loops.

# EXAMPLE

```console
$ buttery -in homer.gif
```

## Features

`-trimStart` / `-trimEnd` drop frames from the far sides of the input animation sequence.

* Trimming animations reduces file size.
* Trimming highlights your favorite motions.
* Trimming also helps to smooth over awkward motions at the start, middle, or end of an animation.
* Trimming can generate unique motion effects, by gluing together animation loops with unexpected timings.

See `buttery -help` for more information.

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/buttery/cmd/buttery@latest
```

# LICENSE

FreeBSD

# RUNTIME REQUIREMENTS

(None)

## Recommended

* [ffmpeg](https://ffmpeg.org/) edits and converts videos
* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's
* [karp](https://github.com/mcandre/karp) for conveniently browsing files and directories
* [mkvtools](https://emmgunn.com/wp/mkvtools-home/) edits MKV videos
* [webm](https://www.webmproject.org/) supports audio in animation loops

# CONTRIBUTING

For more information on developing buttery itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

🧈
