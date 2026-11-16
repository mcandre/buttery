# buttery: a video editor with manual motion smoothing

[![CloudFlare R2 Downloads](https://badgen.net/static/download/Packages/f38020)](#download) [![API Docs](https://badgen.net/static/docs/API/007d9c?icon=go)](https://pkg.go.dev/github.com/mcandre/buttery) [![CI/CD](https://badgen.net/github/checks/mcandre/buttery?icon=github)](https://github.com/mcandre/buttery/actions) [![License](https://badgen.net/static/license/BSD/ab2b29)](LICENSE.md) [![Donate](https://badgen.net/static/sponsor/%E2%99%A5?color=bf3988)](https://github.com/sponsors/mcandre)

# SUMMARY

`buttery` smoothes out transitions at the ends of animated GIFs.

# DOWNLOAD

<table>
  <thead>
    <tr>
      <th>OS</th>
      <th colspan=2>Package</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>macOS 26 Tahoe+</td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.33/macos/buttery-arm64-0.0.33-1.pkg">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.33/macos/buttery-x86_64-0.0.33-1.pkg">Intel</a></td>
    </tr>
    <tr>
      <td>Ubuntu / WSL 2+</td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.33/ubuntu/buttery_0.0.33-1_arm64.deb">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.33/ubuntu/buttery_0.0.33-1_amd64.deb">Intel</a></td>
    </tr>
  </tbody>
</table>

For more platforms and installation methods, see [INSTALL](INSTALL.md).

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
