package gifsicle

import (
	"errors"
	"fmt"
	"image/gif"
	"io"

	"github.com/munchpass/gifsicle-go/embedbinwrapper"
)

// For the -O[level] or --optimize[=level] arguments.
//
// See https://www.lcdf.org/gifsicle/man.html for more information.
type OptimizeLevel string

var (
	// Store only the changed portion of each image. This is the default.
	OPTIMIZE_LEVEL_ONE OptimizeLevel = "1"

	// Store only the changed portion of each image, and use transparency.
	OPTIMIZE_LEVEL_TWO OptimizeLevel = "2"

	// Tries several optimization methods (usually slower, sometimes better results).
	OPTIMIZE_LEVEL_THREE OptimizeLevel = "3"

	// Preserve empty transparent frames (they are dropped by default).
	OPTIMIZE_LEVEL_KEEP_EMPTY OptimizeLevel = "keep-empty"
)

// Gifsicle wraps the gifsicle CLI tool.
type Gifsicle struct {
	binWrapper *embedbinwrapper.EmbedBinWrapper

	inputFile string
	inputGif  *gif.GIF
	input     io.Reader

	outputFile string
	output     io.Writer

	// Integer from 0-200 for GIF compression
	lossy uint

	optimizeLevel OptimizeLevel

	// A number from 2-256 specifying the number of colors (makes a meaningful difference
	// in compressed size)
	numColors *uint
}

func NewGifsicle() (*Gifsicle, error) {
	binWrapper, err := createBinWrapper("gifsicle")
	if err != nil {
		return nil, fmt.Errorf("failed to create binary wrapper for gifsicle: %v", err)
	}

	bin := &Gifsicle{
		binWrapper:    binWrapper,
		optimizeLevel: OPTIMIZE_LEVEL_ONE,
		lossy:         20,
	}

	return bin, nil
}

// Turns on debug mode.
func (g *Gifsicle) Debug() *Gifsicle {
	g.binWrapper = g.binWrapper.Debug()
	return g
}

// InputFile sets image file to convert.
// Input or InputGif called before will be ignored.
func (g *Gifsicle) InputFile(file string) *Gifsicle {
	g.input = nil
	g.inputGif = nil
	g.inputFile = file
	return g
}

// Input sets reader to convert.
// InputFile or InputImage called before will be ignored.
func (g *Gifsicle) Input(reader io.Reader) *Gifsicle {
	g.inputFile = ""
	g.inputGif = nil
	g.input = reader
	return g
}

// InputGif sets gif to convert.
// InputFile or Input called before will be ignored.
func (g *Gifsicle) InputGif(inputGif *gif.GIF) *Gifsicle {
	g.inputFile = ""
	g.input = nil
	g.inputGif = inputGif
	return g
}

// OutputFile specify the name of the output jpeg file.
// Output called before will be ignored.
func (g *Gifsicle) OutputFile(file string) *Gifsicle {
	g.output = nil
	g.outputFile = file
	return g
}

// Output specify writer to write jpeg file content.
// OutputFile called before will be ignored.
func (g *Gifsicle) Output(writer io.Writer) *Gifsicle {
	g.outputFile = ""
	g.output = writer
	return g
}

// For the -O[level] or --optimize[=level] arguments.
//
// See https://www.lcdf.org/gifsicle/man.html for more information.
func (g *Gifsicle) OptimizeLevel(l OptimizeLevel) *Gifsicle {
	g.optimizeLevel = l
	return g
}

// Sets the --lossy parameter.
//
// This parameter ranges from 0-200 (higher value -> more compression).
//
// This is the main compression parameter used in websites like ezgif!
func (g *Gifsicle) Lossy(lossy uint) *Gifsicle {
	if lossy > 200 {
		lossy = 200
	}

	g.lossy = lossy
	return g
}

// Sets the --colors parameter.
//
// This parameter ranges from 2-256 (lower value -> more compression).
func (g *Gifsicle) NumColors(numColors uint) *Gifsicle {
	if numColors < 2 {
		numColors = 2
	}

	if numColors > 256 {
		numColors = 256
	}

	g.numColors = &numColors
	return g
}

// Version returns gifsicle --version
func (g *Gifsicle) Version() (string, error) {
	return version(g.binWrapper)
}

// Reset resets all parameters to default values
func (g *Gifsicle) Reset() *Gifsicle {
	g.lossy = 20
	g.optimizeLevel = OPTIMIZE_LEVEL_ONE
	g.numColors = nil
	return g
}

func (g *Gifsicle) setInput() error {
	if g.input != nil {
		g.binWrapper.StdIn(g.input)
	} else if g.inputGif != nil {
		r, err := createReaderFromGif(g.inputGif)
		if err != nil {
			return fmt.Errorf("createReaderFromGif: %v", err)
		}

		g.binWrapper.StdIn(r)
	} else if g.inputFile != "" {
		g.binWrapper.Arg(g.inputFile)
	} else {
		return errors.New("undefined input")
	}

	return nil
}

func (g *Gifsicle) getOutput() (string, error) {
	if g.output != nil {
		return "", nil
	} else if g.outputFile != "" {
		return g.outputFile, nil
	} else {
		return "", errors.New("undefined output")
	}
}

func (g *Gifsicle) Run() error {
	defer g.binWrapper.Reset()

	g.binWrapper.Arg(fmt.Sprintf("--lossy=%d", g.lossy))
	g.binWrapper.Arg(fmt.Sprintf("--optimize=%s", g.optimizeLevel))

	if g.numColors != nil {
		g.binWrapper.Arg("--colors", fmt.Sprintf("%d", *g.numColors))
	}

	output, err := g.getOutput()
	if err != nil {
		return err
	}

	if output != "" {
		g.binWrapper.Arg("--output", output)
	}

	err = g.setInput()

	if err != nil {
		return err
	}

	if g.output != nil {
		g.binWrapper.SetStdOut(g.output)
	}

	err = g.binWrapper.Run()
	if err != nil {
		return errors.New(err.Error() + ". " + string(g.binWrapper.StdErr()))
	}

	return nil
}
