# buttery: a video editor for smoother GIF loops

![examples/homer.buttery.gif](examples/homer.buttery.gif)

# ABOUT

The popular [GIF](https://en.wikipedia.org/wiki/GIF) format has *technically* supported animation loops for a long time. However, many GIF creators do not take full advantage of the capabilities of the image format. `buttery` helps to edit GIF files into smoother animation sequences.

For example, some GIF's may consist of a single frame. Some GIF's may animate for a few seconds but then terminate. Still other GIF's may loop endlessly, but with jerky motion along the sequence. `buttery` can help resolve many of these quirks.

# EXAMPLE

```console
$ cd examples

$ buttery -in homer.gif

$ ls
homer.gif
homer.buttery.gif
```

## Common options

`-trimStart` / `-trimEnd` can drop frames from sides of the input GIF sequence. Trimming is useful for cutting down long animations to a specific subsequence. Trimming may facilitate motion smoothing. When used artfully, trimming can even generate unique motion.

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
