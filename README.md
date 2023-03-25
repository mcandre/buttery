# buttery: a GIF editor for generating smoother animation loops

![examples/homer.buttery.gif](examples/homer.buttery.gif)

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

* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's.

🧈
