# buttery: a video editor with manual motion smoothing

[![Docker Pulls](https://img.shields.io/docker/pulls/n4jm4/buttery)](https://hub.docker.com/r/n4jm4/buttery) [![Donate](https://img.shields.io/badge/GUMROAD-36a9ae?style=flat&logo=gumroad&logoColor=white)](https://mcandre.gumroad.com/)

![examples/homer.buttery.gif](examples/homer.buttery.gif)

![examples/cinnamoroll.buttery.gif](examples/cinnamoroll.buttery.gif)

# ABOUT

`buttery` generates continuous GIF loops.

# EXAMPLES

```console
$ cd examples

$ buttery homer.gif

$ buttery -transparent -stitch FlipH cinnamoroll.gif
```

See `buttery -help` for more options.

# API DOCUMENTATION

https://pkg.go.dev/github.com/mcandre/buttery

# INSTALLATION

See [INSTALL.md](INSTALL.md).

# LICENSE

BSD-2-Clause

# LINT

`buttery -check <GIF>` can act as a linter for basic GIF file format integrity. In the event of a corrupt GIF file, the program emits a brief message and exits non-zero.

# INSPECT

`buttery -getFrames <GIF>` reports the frame count. This is useful for planning edits, particularly towards the far end of the original animation sequence.

# BACKGROUNDS

The `-transparent` option changes the disposal mode from none to background, and changes the background from black to clear.

# TRANSITIONS

## Mirror

The `-stitch Mirror` option is the primary loop smoothing transition, and the default transition setting.

Mirror twists the GIF timeline around like a Mobius strip, so that it arrives naturally back at the start. This is useful for smoothing GIF's that present misaligned images at the extreme ends of the loop.

We can diagram logically how Mirror works, by examining its effect on the frame sequence. With the notation:

```text
real frame sequence (successive sequence repeated during infinite loop playback...)
```

### Before

```text
1 2 3 (1 2 3 ...)
```

Each restart of the loop has a jarring visual jump from frame 3 to its successor frame 1.

### After

```text
1 2 3 2 (1 2 3 2 ...)
```

Of course, running the `buttery` editor yourself is the best way to appreciate how it works.

By mirroring the sequence backward in time, we remove the biggest visual jump. The overall visual effect is that of a sailor rowing back and forth in place. The Mirror transition often dramatically improves the smoothness of a GIF loop.

However, some motion may still appear awkward with mirroring, such as sharp, quick motions towards the extreme ends of the loop, or motions that appear to defy physical entropy. For this reason, we provide alternative transitions and other editing tools, described below.

## FlipH / FlipV

The transision settings `-stitch FlipH` or `-stitch FlipV` disguise jarring misalignment, by reflecting the frames horizontally or vertically.

With the notation:

* `R`: An original "right" frame
* `L`: A frame reflected horizontally "leftward"
* `U`: A original "upright" frame
* `D`: A frame reflected vertically "downward"

### Before

FlipH:

```text
R R R (R R R ...)
```

FlipV:

```text
U U U (U U U ...)
```

### After

FlipH:

```text
R R R L L L (R R R L L L ...)
```

FlipV:

```text
U U U D D D (U U U D D D ...)
```

The `FlipH`/`FlipV` transitions are snappy, with an effect like rebounding a tennis ball across a net.

## Shuffle

The `-stitch Shuffle` transition setting randomizes the frame sequence.

### Before

```text
1 2 3 4 5 6
```

### After

Example ordering:

```text
6 3 1 4 2 5
```

Naturally, the more unique frames available, the more opportunity for unique random orderings.

This transition tends to artificially accelerate the perceived animation speed.

This transition hides a single jarring misalignment, in the noise of a completely random, spastic animation.

## PanH / PanV

The `PanH` / `PanV` transitions offset the canvas at `-panVelocity <n>` pixels per frame.

## Fade

The transision setting `-stitch Fade` applies fade to black, fade to white, etc. time color gradient effects.

### Before

```text
no_fade ... no_fade ... no_fade ... no_fade ... no_fade
```

### After

```text
max_fade ... less_fade ... no_fade ... less_fade ... max_fade
```

`-fadeColor 0xRRGGBB` customizes the fade hue (default: black).

`-fadeRate <v>` adjusts fade velocity (default: 1.0).

## None

The `-stitch None` transition setting applies no particular transition at all between animation cycles. In art, sometimes less is more.

# SUPERCUTS

Animation smoothing takes a long time. We recommend pre-cutting your source assets to the desired subsequence. Every frame removed from the input GIF makes the `buttery` editing process faster.

Often, animations appear to accelerate when frame are removed. This is not always a bad thing; sometimes a fast animation helps to smooth over more subtle details.

## Trim Start / End

The `-trimStart <n>` / `-trimEnd <n>` options drop `n` frames from the start and/or end of the original sequence. Zero indicates no trimming.

For brevity, we will now assume the None transition and elide the successive sequence repetitions.

### Before

```text
1 2 3 4 5
```

### After

With `-trimStart 1`:

```text
2 3 4 5
```

With `-trimEnd 1`:

```text
1 2 3 4
```

With `-trimStart 1` and `-trimEnd 1`:

```text
2 3 4
```

## Trim Edges

For convenience, we provide a similar option `-trimEdges <n>`. This drops `n` frames from both sides of the original sequence. Zero indicates no trimming.

### Before

```text
1 2 3 4 5
```

### After

With `-trimEdges 1`:

```text
2 3 4
```

## Window

The `-window <n>` option truncates the original sequence to a fixed frame count. This is helpful for cutting down long animations. Zero indicates no truncation.

### Before

```text
1 2 3 4 5
```

### After

With `-window 3`:

```text
1 2 3
```

With `-window 3` and `-trimStart 1`:

```text
2 3 4
```

## Cut Interval

The `-cutInterval <n>` option removes every nth frame from the original sequence.

This can mitigate some oscillation, such as lighting fluctuations from fans.

### Before

```text
1 2 3 4 5 6 7 8
```

### After

With `-cutInterval 2`:

```text
1 3 5 7
```

This can also artificially accelerate the perceived speed of the animation. Useful when want to accelerate an animation already scaled down to 2cs per frame.

# SHIFT

The `-shift <offset>` option performs a circular, leftward shift on the original sequence. This is useful for fine tuning how the GIF's very first cycle presents, before entering successive loops.

Zero is the neutral shift. A negative offset indicates rightward shift.

### Before

```text
1 2 3
```

### After

With `-shift 1`:

```text
2 3 1
```

With `-shift -1`:

```text
3 1 2
```

# SCALE DELAY

The `-scaleDelay <factor>` option adjusts animation speed, by multiplying each frame delay by the given factor.

`1` = `1.0` is the neutral, and default factor.

Negative values reverse the original sequence.

We can diagram this in terms of the frame delays, expressed in centiseconds. That is, `4cs` indicates 4 centisec = 4/100 sec between advancing to the next frame.

### Before

```text
4cs 6cs 8cs
```

### After

With `-scaleDelay 2`:

```text
8cs 12cs 16cs
```

With `-scaleDelay 0.5`:

```text
2cs 3cs 4cs
```

With `-scaleDelay -1`:

```text
8cs 6cs 4cs
```

For compatibility with a wide range of GIF viewers, the resulting delay is upheld to a lower bound of 2cs.

# LOOP COUNT

The `-loopCount <n>` option configures the low-level GIF loop counter setting. According to the GIF standard:

* `-1` indicates loop exactly once.
* `0` indicates infinite, endless looping (default).
* `n` indicates n replays after the first play = 1 + n total iterations.

### Before

```text
1 2 3
```

### After

With `-loopCount 0`:

```text
1 2 3 (1 2 3 ...)
```

# SEE ALSO

* [ffmpeg](https://ffmpeg.org/) edits and converts videos
* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's
* [harmonica](https://github.com/mcandre/harmonica) repackages comic archives
* [ImageMagick](https://imagemagick.org/) converts between multimedia formats, including GIF and WEBP
* [mkvtools](https://emmgunn.com/wp/mkvtools-home/) edits MKV videos
* [tigris](https://github.com/mcandre/tigris) provides (Kindle) comic book archival utilities
* [VLC](https://www.videolan.org/vlc/) plays numerous multimedia formats
* [webm](https://www.webmproject.org/) supports audio in animation loops

🧈
