# buttery: a video editor with manual motion smoothing

[![CloudFlare R2 install media downloads](https://img.shields.io/badge/Packages-F38020?logo=Cloudflare&logoColor=white)](#download) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/buttery) [![Test](https://github.com/mcandre/buttery/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/buttery/actions/workflows/test.yml) [![license](https://img.shields.io/badge/license-BSD-0)](LICENSE.md) [![Donate](https://img.shields.io/badge/%E2%99%A5-Sponsor-BF3988)](https://github.com/sponsors/mcandre)

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
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.29/macos/buttery-arm64-0.0.29-1.pkg">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.29/macos/buttery-x86_64-0.0.29-1.pkg">Intel</a></td>
    </tr>
    <tr>
      <td>Ubuntu / WSL 2+</td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.29/ubuntu/buttery_0.0.29-1_arm64.deb">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.29/ubuntu/buttery_0.0.29-1_amd64.deb">Intel</a></td>
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
buttery -stitch Fade -fadeColor '0xCC5500' red-panda.gif
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
* [ImageMagick](https://imagemagick.org/) converts between multimedia formats, including GIF and WEBP
* [mcandre/harmonica](https://github.com/mcandre/harmonica) repackages comic archives
* [mcandre/tigris](https://github.com/mcandre/tigris) provides (Kindle) comic book archival utilities
* [mkvtools](https://emmgunn.com/wp/mkvtools-home/) edits MKV videos
* [VLC](https://www.videolan.org/vlc/) plays numerous multimedia formats
* [webm](https://www.webmproject.org/) supports audio in animation loops

🧈
