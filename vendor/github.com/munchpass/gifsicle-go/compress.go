package gifsicle

import (
	"fmt"
	"image/gif"
	"io"
)

type Options struct {
	// This parameter ranges from 0-200 (higher value -> more compression).
	//
	// This is the main compression parameter used in websites like ezgif!
	Lossy uint

	// Additional compression optimization.
	OptimizeLevel OptimizeLevel

	// Optional parameter to specify the number of colors (2-256) for the GIF
	//
	// Specifying less colors compresses the GIF more.
	// If this is 0, gifsicle will default to using the image's color map.
	NumColors uint
}

// Shortcut function to compress GIFs quickly and easily with gifsicle.
func Compress(w io.Writer, g *gif.GIF, o *Options) error {
	gifsicleCli, err := NewGifsicle()
	if err != nil {
		return fmt.Errorf("NewGifsicle failed: %v", err)
	}

	if o != nil {
		gifsicleCli.Lossy(o.Lossy)
		gifsicleCli.OptimizeLevel(o.OptimizeLevel)
		if o.NumColors >= 2 {
			gifsicleCli.NumColors(o.NumColors)
		}
	}

	return gifsicleCli.InputGif(g).Output(w).Run()
}

/*
Shortcut function to compress GIFs quickly and easily with gifsicle
using an io.Reader like a file or buffer.

Example:

	testFilePath := path.Join("testfiles", "portrait_3mb.gif")
	testFile, err := os.Open(testFilePath)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = gifsicle.CompressFromReader(&buf, testFile, &gifsicle.Options{
		Lossy:         200,
		OptimizeLevel: gifsicle.OPTIMIZE_LEVEL_THREE,
		NumColors:     256,
	})
	if err != nil {
		return err
	}
*/
func CompressFromReader(w io.Writer, r io.Reader, o *Options) error {
	gifsicleCli, err := NewGifsicle()
	if err != nil {
		return fmt.Errorf("NewGifsicle failed: %v", err)
	}

	if o != nil {
		gifsicleCli.Lossy(o.Lossy)
		gifsicleCli.OptimizeLevel(o.OptimizeLevel)
		if o.NumColors >= 2 {
			gifsicleCli.NumColors(o.NumColors)
		}
	}

	return gifsicleCli.Input(r).Output(w).Run()
}
