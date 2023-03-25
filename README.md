# buttery: a GIF looper

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

`-trimStart <frames>`: drop a number of frames from the start of the loop. Useful for cleaning up long animations.

`-trimEnd <frames>`: drop a number of frames from the end of the loop. Useful for generating smoother loops.

See `buttery -help` for more detail.

# LICENSE

FreeBSD

# RUNTIME REQUIREMENTS

(None)

## Recommended

* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's.

🧈
