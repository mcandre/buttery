# buttery: a video editor with manual motion smoothing

![examples/homer.buttery.gif](examples/homer.buttery.gif)

# ABOUT

`buttery` generates continuous GIF loops.

# EXAMPLE

```console
$ buttery homer.gif
```

See `buttery -help` for more options.

# DOWNLOAD

https://github.com/mcandre/buttery/releases

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/buttery/cmd/buttery@latest
```

# API DOCUMENTATION

https://pkg.go.dev/github.com/mcandre/buttery

# LICENSE

FreeBSD

# RUNTIME REQUIREMENTS

(None)

# CONTRIBUTING

For more information on developing buttery itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# COMMON FEATURES

`-trimStart` / `-trimEnd` drop frames from the far sides of the input animation sequence.

* Trimming animations reduces file size.
* Trimming highlights your favorite motions.
* Trimming also helps to smooth over awkward motions at the start, middle, or end of an animation.
* Trimming can generate creative motion effects, by gluing together animation loops at serendipitious frame timings.

# GENERAL TIPS

## Transitions

`-stitch Mirror` is the default continuity strategy. It works like an accordion, following the incoming sequence by replaying the incoming sequence backwards to form a loop.

`-stitch FlipH` / `-stitch FlipV` follow the incoming sequence by reflecting the sequence horizontally or vertically.

`-stitch None` presents the sequence with no special continuity transition. This is mainly useful for simply enabling the infinite loop setting on a GIF.

## Supercuts

Animation smoothing takes a long time. Pre-cut your GIF's to highlight the specific desired sequence, for faster overall editing.

The ends of a loop, and the middle of a mirrored loop, are highly sensitive to stuttery motion in the opposite direction of adjacent frames. Often, as you `-trimStart` / `-trimEnd`, the continuity appears to worsen right up until the critical frames are removed. You can speed up your editing workflow with binary search: Instead of incrementing or decrementing values by one, try doubling, quadrupling, halving, quartering, etc. Experiment.

Motion appears to accelerate with fewer frames. This is not always a bad thing; sometimes a fast animation helps to smooth over more subtle details.

## Speed

The maximum theoretical GIF speed is 100 FPS, though in practice many GIF viewers, such as Web browsers, may support slower speeds, such as 50 FPS or lower. Worse, some GIF viewers interpret a GIF frame of 0.01 sec fast delay, as a reset to 1 sec slow delay.

Generally stick to `-speed` factors inside of `0.1` (quite slow) and `3.0` (quite fast).

Note that some GIF's are already using the quickest delay setting possible. In that case, a `-speed` greater than or equal to `1`, may not alter the animation speed.

# SEE ALSO

* [ffmpeg](https://ffmpeg.org/) edits and converts videos
* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's
* [ImageMagick](https://imagemagick.org/) converts between multimedia formats, including GIF and WEBP
* [karp](https://github.com/mcandre/karp) for conveniently browsing files and directories
* [mkvtools](https://emmgunn.com/wp/mkvtools-home/) edits MKV videos
* [VLC](https://www.videolan.org/vlc/) plays numerous multimedia formats
* [webm](https://www.webmproject.org/) supports audio in animation loops

🧈
