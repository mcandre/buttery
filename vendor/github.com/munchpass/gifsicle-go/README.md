# gifsicle-go <!-- omit in toc -->

A Go wrapper around [`gifsicle`](https://www.lcdf.org/gifsicle/man.html) with bundled binaries for quick and easy GIF compression with just a single `go get`!

`gifsicle` is one of the most popular GIF compression tools used by many online GIF compression websites like [ezgif.com](https://ezgif.com/).

## Table of Contents <!-- omit in toc -->

- [Install](#install)
- [Usage](#usage)
  - [`Compress`](#compress)
  - [`Gifsicle`](#gifsicle)
- [Contributing](#contributing)

## Install

We bundle the binaries through embeds, so all you need is a `go get`:

```bash
go get github.com/munchpass/gifsicle-go
```

No need to have `gifsicle` pre-installed!

## Usage

### `Compress`

If you just want to quickly compress a GIF and write the result into a buffer:

```go
	var buf bytes.Buffer
	err = gifsicle.Compress(&buf, decodedGif, &gifsicle.Options{
		Lossy:         80,
		OptimizeLevel: gifsicle.OPTIMIZE_LEVEL_THREE,
	})
	if err != nil {
		return err
	}
```

### `Gifsicle`

This library also provides the `Gifsicle` struct which is a thin wrapper around the `gifsicle` CLI tool.

Here are some examples:

Compress an input GIF file and write the result to an output GIF file:

```go
	sourceGif := path.Join("testfiles", "portrait_3mb.gif")
	outputGif := path.Join("testoutput", "portrait_3mb_output.gif")
	err = g.InputFile(sourceGif).
		OutputFile(outputGif).
		Lossy(80).
		OptimizeLevel(gifsicle.OPTIMIZE_LEVEL_THREE).
		Run()
```

Compress an input GIF object and write the result to an output GIF file:

```go
	err = g.InputGif(sourceGif).
		OutputFile(outputGif).
		Lossy(80).
		OptimizeLevel(gifsicle.OPTIMIZE_LEVEL_THREE).
		Run()
```

Compress an input GIF object and write the result to an output buffer:

```go
    var buf bytes.Buffer
	err = g.InputGif(sourceGif).
		Output(&buf).
		Lossy(80).
		OptimizeLevel(gifsicle.OPTIMIZE_LEVEL_THREE).
		Run()
```

## Contributing

This library is still in its early stage, but is actively used in production. We welcome any PRs/suggestions!
