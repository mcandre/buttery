# buttery: a video editor with manual motion smoothing

![examples/homer.buttery.gif](examples/homer.buttery.gif)

# ABOUT

`buttery` generates continuous GIF loops.

# EXAMPLE

```console
$ buttery homer.gif
```

See `buttery -help` for more options.

# DOWNLOAD

https://github.com/mcandre/buttery/releases

# INSTALL FROM SOURCE

```console
$ go install github.com/mcandre/buttery/cmd/buttery@latest
```

# API DOCUMENTATION

https://pkg.go.dev/github.com/mcandre/buttery

# LICENSE

FreeBSD

# RUNTIME REQUIREMENTS

(None)

# CONTRIBUTING

For more information on developing buttery itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# LINT

`buttery -check <GIF>` can act as a linter for basic GIF file format integrity. In the event of a corrupt GIF file, the program emits a brief message and exits non-zero.

# INSPECT

`buttery -getFrames <GIF>` reports the frame count. This is useful for planning edits, particularly towards the far end of the original animation sequence.

# TRANSITIONS

## Mirror

`-stitch Mirror` is the primary transition smoothing option, and the default `buttery` behavior.

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

By mirroring the sequence backward in time, we remove the biggest visual jump. The overall visual effect is that of a sailor rowing back and forth in place. The Mirror transition can dramatically improve the smoothness of a GIF loop, and it's our default, primary editing tool.

Mirroring may be slightly more compact in terms of frame count and file size, than FlipH or FlipV.

However, some motion may still appear awkward with mirroring, such as sharp, quick motions towards the extreme ends of the loop, or motions that appear to defy physical entropy. For this reason, we provide alternative transitions and other editing tools.

## FlipH / FlipV

The `-stitch FlipH` or `-stitch FlipV` transition disguises jarring misalignment by reflecting the frames about an axis.

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

This transition is more playful, with an effect akin to rebounding a tennis ball over a net.

## None

The `-stitch None` transition applies no transition at all between animation cycles. In art, sometimes less is more.

### Before

```text
1 2 3 (1 2 3 ...)
```

### After

```text
1 2 3 (1 2 3 ...)
```

In other words, this option disables the other Mirror, FlipH, and FlipV transitions.

# SUPERCUTS

Animation smoothing takes a long time. We recommend pre-cutting your source assets to the desired subsequence. Every frame removed from the input GIF makes the `buttery` editing process faster.

Often, animations appear to accelerate when frame are removed. This is not always a bad thing; sometimes a fast animation helps to smooth over more subtle details.

## Trim Start / End

The `-trimStart <n>` / `-trimEnd <n>` options drop `n` frames from the start and/or end of the original sequence.

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

For convenience, we provide a similar option `-trimEdges <n>`. This drops `n` frames from both sides of the original sequence.

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

The `-window <n>` option truncates the original sequence to a fixed frame count. This is helpful for cutting down long animations.

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

# SHIFT

The `-shift <offset>` option performs a circular, leftward shift on the original sequence. This is useful for fine tuning how the GIF's very first cycle presents, before entering successive loops.

A negative offset indicates rightward shift. Zero is the neutral shift.

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

# REVERSE

The `-reverse` option reorders the original sequence backwards.

### Before

```text
1 2 3
```

### After

```text:
3 2 1
```

# SPEED

The `-speed <factor>` option adjusts animation speed. Speed is expressed as a factor relative to the original GIF frame delay. We recommend using values in the range `0.2` (slow) to `2.0` (fast).

We can diagram this in terms of the frame delays, expressed in centiseconds. That is, `4cs` indicates 4 centisec = 4/100 sec between advancing to the next frame.

### Before

```text
4cs 4cs 4cs
```

### After

With `-speed 2.0`:

```text
2cs 2cs 2cs
```

With `-speed 0.5`:

```text
8cs 8cs 8cs
```

`1.0` is the neutral factor. Speed cannot be zero or negative.

### Quirks, Quickly

GIFs have some further quirks worth noting, regarding animation speed.

The lowest delay value is zero, though in GIF format speak, that indicates a reset to a default 1 sec delay, which is slow.

In theory, the fastest GIF delay is 1 centisec. However, this equates to an FPS frame rate higher than most computer systems can handle. In practice, the fastest GIF delay is 2 centisec, nearly 60 FPS. Unfortunately, some GIF rendering engines such as Google Chrome, may also treat 1 centisec as a reset to default 1 sec delay, which is slow. In order to ensure wide compatibility with many different GIF viewing applications, `buttery` enforces a lower bound of 2 centisec for frame delays.

Frame delays are sensitive to high speed factors. Even most factors like `3.0` and higher, may produce a stuttery or blurred effect.

Short, single digit frame delays are even more sensitive to speed factors. This is due to inherent limitations in delay integer precision. As a result, many different values of factors can yield animations with identical animation speeds.

Note that some GIF's are already using the quickest delay setting possible, about 2cs. In that case, a speed factor greater than or equal to `1`, may not have the desired effect.

# SEE ALSO

* [ffmpeg](https://ffmpeg.org/) edits and converts videos
* [gifenc.sh](https://github.com/thevangelist/FFMPEG-gif-script-for-bash) converts numerous video formats to animated GIF's
* [ImageMagick](https://imagemagick.org/) converts between multimedia formats, including GIF and WEBP
* [karp](https://github.com/mcandre/karp) for conveniently browsing files and directories
* [mkvtools](https://emmgunn.com/wp/mkvtools-home/) edits MKV videos
* [VLC](https://www.videolan.org/vlc/) plays numerous multimedia formats
* [webm](https://www.webmproject.org/) supports audio in animation loops

🧈
