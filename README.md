# buttery: a video editor for smoother animations

![examples/homer.buttery.gif](examples/homer.buttery.gif)

# ABOUT

Artists have enjoyed creating digital animations for many years, thanks to popular file formats like the legendary [GIF](https://en.wikipedia.org/wiki/GIF). Woohoo!

File formats as a *technical subject* turn out to offer interesting tradeoffs, including errors as well as amazing features. `buttery` is here to help artists navigate these tradeoffs. We want you to enjoy the full power of modern, programmable editing technology.

For example, browsing assorted GIF's online reveals:

* Jerky motion
* File corruption
* Animations too long or too short
* GIF's consisting of a single frame
* Animations that halt after a few seconds

`buttery` treat many of these issues with simple, programmable commands. Our goals include making animation easier and more *expressive*, so that you can focus less on technicals and more on showcasing your work.

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
