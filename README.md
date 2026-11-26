# buttery: a video editor with manual motion smoothing

[![CloudFlare R2 Downloads](https://badgen.net/static/download/Packages/f38020?icon=cloudflare)](https://mcandre.github.io/buttery/) [![API Docs](https://badgen.net/static/docs/API/007d9c?icon=go)](https://pkg.go.dev/github.com/mcandre/buttery) [![CI/CD](https://badgen.net/github/checks/mcandre/buttery?icon=github)](https://github.com/mcandre/buttery/actions) [![License](https://badgen.net/static/license/BSD/ab2b29?icon=freebsd)](LICENSE.md) [![AI](https://badgen.net/static/ai/Gemini|hint?color=886fbf&icon=google-gemini)](AI-DECLARATION.md) [![Donate](https://badgen.net/static/sponsor/%E2%99%A5?color=bf3988&icon=github)](https://github.com/sponsors/mcandre)

![stick 'o butter man](buttery.png)

# SUMMARY

`buttery` smoothes out transitions at the ends of animated GIFs.

For details on tuning buttery, see [CONFIGURATION](CONFIGURATION.md).

For details on building from source, see [DEVELOPMENT](DEVELOPMENT.md).

# EXAMPLES

Found in the [examples/](examples) directory.

## Default Loop Stitch

By default, buttery smoothes out animations by copying frames in reverse. As well, the GIF loop count flag is set to infinite.

Homer changes from popping instantaneously out of a bush, to emerging back out of a bush.

| Before | After |
| :----: | :---: |
| ![examples/homer.gif](examples/homer.gif) | ![examples/homer.buttery.gif](examples/homer.buttery.gif) |

```sh
buttery homer.gif
```

## Color Fading

The classic fade-to-black (or orange!) transition. Dream-like.

| Before | After |
| :----: | :---: |
| ![examples/red-panda.gif](examples/red-panda.gif) | ![examples/red-panda.buttery.gif](examples/red-panda.buttery.gif) |

```sh
buttery -stitch Fade -fadeColor "#CC5500" red-panda.gif
```

## FlipH/FlipV

Reflecting frames around a horizontal or vertical axis, starting at the animation halfway point in time, is another great way to smooth out transitions.

| Before | After |
| :----: | :---: |
| ![examples/cinnamoroll.gif](examples/cinnamoroll.gif) | ![examples/cinnamoroll.buttery.gif](examples/cinnamoroll.buttery.gif) |

```sh
buttery -transparent -stitch FlipH cinnamoroll.gif
```

Run `buttery -help` for even more options.

# RESOURCES

Prior art and personal plugs.

* [ffmpeg](https://ffmpeg.org/) edits and converts videos
* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's
* [ImageMagick](https://imagemagick.org/#gsc.tab=0) converts between multimedia formats, including GIF and WEBP
* [mcandre/harmonica](https://github.com/mcandre/harmonica) repackages comic archives
* [mcandre/tigris](https://github.com/mcandre/tigris) provides (Kindle) comic book archival utilities
* [mkvtools](https://emmgunn.com/wp/mkvtools-home/) edits MKV videos
* [VLC](https://www.videolan.org/vlc/) plays numerous multimedia formats
* [webm](https://www.webmproject.org/) supports audio in animation loops

🧈
