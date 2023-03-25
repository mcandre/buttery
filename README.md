# buttery: a video editor for smoother GIF loops

![examples/homer.buttery.gif](examples/homer.buttery.gif)

# ABOUT

Animators have enjoyed creating digital animations for many years, thanks to the popular [GIF](https://en.wikipedia.org/wiki/GIF) format. In fact, GIF as a *technical* subject has not been the easiest for artists to consume.

Take a tour of available GIF's online, and you may find:

* File corruption
* Very long animations
* Animations too long or too short
* Animations that halt after a few seconds
* GIF's consisting of a single frame
* Jerky motion

`buttery` treat many of these issues with simple, programmable edits. Our goal is to make the GIF format manipulation process easier, and to generate more pleasing images. Showcase your craft!

# EXAMPLE

```console
$ cd examples

$ buttery -in homer.gif

$ ls
homer.gif
homer.buttery.gif
```

## Common options

`-trimStart` / `-trimEnd` can drop frames from sides of the input GIF sequence. Trimming is useful for cutting down long animations to a specific subsequence.

This can facilitate motion smoothing even further. When used artfully, the simple act of trimming frames can even generate unique motion effects.

Pro tip: When cutting very long sequences, start by trimming the end first, which is often a faster operation than trimming the start. This is due to the inner workings of GIF's compression model.

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
