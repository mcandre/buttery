# buttery: a video editor for smoother GIF loops

![examples/homer.buttery.gif](examples/homer.buttery.gif)

# ABOUT

The popular [GIF](https://en.wikipedia.org/wiki/GIF) format has technically supported infinitely looping animations for a long time. However, many GIF images do not take advantage of the full capabilities of this image format.

For example, some GIF's may consist of a single frame. Some GIF's may animate for a few seconds but then terminate. Still other GIF's may loop endlessly, but with jerky motion. `buttery` helps to edit these sequences into smoother, more pleasing animations.

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

* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's
* [karp](https://github.com/mcandre/karp) for conveniently browsing files and directories

🧈
