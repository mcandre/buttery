# buttery: a video editor with manual motion smoothing

![examples/homer.buttery.gif](examples/homer.buttery.gif)

# ABOUT

`buttery` generates continuous GIF loops.

# EXAMPLE

```console
$ buttery -in homer.gif
```

# DOWNLOAD

https://github.com/mcandre/buttery/releases

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/buttery/cmd/buttery@latest
```

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

`-mirror=false` disables accordian mirroring.

See `buttery -help` for more information.

# GENERAL TIPS

## Supercuts

Animation smoothing takes a long time. Pre-cut your GIF's to highlight the specific desired sequence, for faster overall editing.

The ends of a loop, and the middle of a mirrored loop, are highly sensitive to stuttery motion in the opposite direction of adjacent frames. Often, as you `-trimStart` / `-trimEnd`, the continuity appears to worsen right up until the critical frames are removed. You can speed up your editing workflow with binary search: Instead of incrementing or decrementing values by one, try doubling, quadrupling, halving, quartering, etc. Experiment.

Motion appears to accelerate with fewer frames. This is not always a bad thing; sometimes a fast animation helps to smooth over more subtle details.

## Speed

The animation `-speed` factor is also fairly sensitive. For example, a loop may feel much faster at `1.5` (x) speed compared to the default `1` (x) speed.

The maximum theoretical GIF speed is 100 FPS, though in practice many GIF viewers, such as Web browsers, may support slower speeds, such as 50 FPS or lower. Worse, some GIF viewers interpret a GIF frame of 0.01 sec fast delay, as a reset to 1 sec slow delay. Generally stick to `-speed` factors between `0.1` (slow) to `3.0` (fast).

# SEE ALSO

* [ffmpeg](https://ffmpeg.org/) edits and converts videos
* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's
* [karp](https://github.com/mcandre/karp) for conveniently browsing files and directories
* [mkvtools](https://emmgunn.com/wp/mkvtools-home/) edits MKV videos
* [webm](https://www.webmproject.org/) supports audio in animation loops

🧈
