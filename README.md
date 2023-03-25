# buttery: a video editor with motion smoothing

![examples/homer.buttery.gif](examples/homer.buttery.gif)

# EXAMPLE

`buttery` generates smoothe GIF loops.

```console
$ cd examples

$ buttery -in homer.gif

$ ls
homer.gif
homer.buttery.gif
```

## Common options

`buttery` includes additional options for editing animation files.

## Trimming

`-trimStart` / `-trimEnd` drop frames from either side of the input animation sequence.

* Trimming animation sequences reduces file size.
* Trimming cuts down time from quite long animations, to highlight a specific subsequence.
* Trimming also helps to smooth over awkward motions that may occur at the start, middle, or end of an animation.
* Trimming can generate unique motion effects, by gluing together animation loops at unexpected timings.

See `buttery -help` for more information.

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

🧈
