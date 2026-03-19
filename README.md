# buttery: a video editor with manual motion smoothing

[![CloudFlare R2 install media downloads](https://img.shields.io/badge/Packages-F38020?logo=Cloudflare&logoColor=white)](#download) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/mcandre/buttery) [![Test](https://github.com/mcandre/buttery/actions/workflows/test.yml/badge.svg)](https://github.com/mcandre/buttery/actions/workflows/test.yml) [![license](https://img.shields.io/badge/license-BSD-0)](LICENSE.md) [![Donate](https://img.shields.io/badge/%E2%99%A5-Sponsor-BF3988)](https://github.com/sponsors/mcandre)

![examples/cinnamoroll.buttery.gif](examples/cinnamoroll.buttery.gif)

# ABOUT

`buttery` generates continuous GIF loops.

# EXAMPLES

```console
% cd examples

% buttery homer.gif

% buttery -transparent -stitch FlipH cinnamoroll.gif
```

For more CLI options, run `buttery -help`.

For practical usage information, see our [usage guide](USAGE.md).

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
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.27/macos/buttery-arm64-0.0.27-1.pkg">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.27/macos/buttery-x86_64-0.0.27-1.pkg">Intel</a></td>
    </tr>
    <tr>
      <td>Ubuntu / WSL 2+</td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.27/ubuntu/buttery_0.0.27-1_arm64.deb">ARM</a></td>
      <td><a href="https://pub-d141861718d342d19cfd516f2569755e.r2.dev/buttery-0.0.27/ubuntu/buttery_0.0.27-1_amd64.deb">Intel</a></td>
    </tr>
  </tbody>
</table>

For more platforms and installation methods, see our [install guide](INSTALL.md).

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
