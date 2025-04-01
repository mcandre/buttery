package buttery

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"os"

	"github.com/andybons/gogif"
	"github.com/anthonynsimon/bild/transform"
)

// Config models a set of animation editing manipulations.
type Config struct {
	// Transparent preserves clear animations (Default false).
	Transparent bool

	// TrimEdges removes frames from the start and end of the incoming sequence (Default zero).
	TrimEdges int

	// TrimStart removes frames from the start of the incoming sequence (Default zero).
	TrimStart int

	// TrimEnd removes fromes frames from the end of the incoming sequence (Default zero).
	TrimEnd int

	// CutInterval removes every nth frame from the incoming sequence (Default zero).
	CutInterval int

	// Window truncates frames from the incoming sequence (Default zero).
	//
	// Zero indicates no window truncation.
	Window int

	// Shift moves the start of the sequence leftward (Default zero).
	Shift int

	// Stitch denotes a loop continuity transition (Default Mirror).
	Stitch Stitch

	// ScaleDelay multiplies each frame delay by a factor (Default 1.0).
	//
	// The resulting delay is upheld to a lower bound of 2 centisec.
	//
	// A negative scale delay reverses the incoming sequence.
	ScaleDelay float64

	// LoopCount denotes how many times to play the animation (Default 0).
	//
	// -1 indicates one play.
	// 0 indicates infinite, endless plays.
	// N indicates 1+N iterations.
	LoopCount int
}

// NewConfig generates a default Config.
func NewConfig() Config {
	return Config{
		Stitch:     Mirror,
		ScaleDelay: 1.0,
	}
}

// Validate checks for basic Config integrity.
func (o *Config) Validate() error {
	if o.TrimEdges < 0 {
		return errors.New("trim edges cannot be negative")
	}

	if o.TrimStart < 0 {
		return errors.New("trim start cannot be negative")
	}

	if o.TrimEnd < 0 {
		return errors.New("trim end cannot be negative")
	}

	if o.CutInterval < 0 || o.CutInterval == 1 {
		return errors.New("cut interval cannot be less than two")
	}

	if o.Window < 0 {
		return errors.New("window cannot be negative")
	}

	return o.Stitch.Validate()
}

// Edit applies the configured GIF manipulations.
func (o *Config) Edit(destPth string, sourceGif *gif.GIF) error {
	sourcePaletteds := sourceGif.Image
	sourcePalettedsLen := len(sourcePaletteds)

	if o.TrimStart+o.TrimEnd >= sourcePalettedsLen {
		return errors.New("minimum 1 output frame")
	}

	var reverse bool
	scaleDelay := o.ScaleDelay

	if scaleDelay < 0 {
		reverse = true
		scaleDelay *= -1.0
	}

	window := o.Window

	if window > sourcePalettedsLen-o.TrimStart-o.TrimEnd {
		return errors.New("window longer than subsequence")
	}

	sourceDelays := sourceGif.Delay
	sourceWidth, sourceHeight := GetDimensions(sourcePaletteds)
	canvasImage := image.NewRGBA(image.Rect(0, 0, sourceWidth, sourceHeight))
	canvasBounds := canvasImage.Bounds()
	paletteSize := GetPaletteSize(sourcePaletteds)
	clonePaletteds := make([]*image.Paletted, sourcePalettedsLen)
	quantizer := gogif.MedianCutQuantizer{NumColor: paletteSize}
	var disposals []byte
	c := color.Alpha16{0}

	if o.Transparent {
		c = color.Transparent
	}

	draw.Src.Draw(canvasImage, canvasBounds, &image.Uniform{sourcePaletteds[0].Palette.Convert(c)}, image.Point{})

	for i, sourcePaletted := range sourcePaletteds {
		im := canvasImage

		if o.Transparent {
			im = image.NewRGBA(image.Rect(0, 0, sourceWidth, sourceHeight))
		}

		draw.Over.Draw(im, canvasBounds, sourcePaletted, image.Point{})
		clonePaletted := image.NewPaletted(canvasBounds, sourcePaletted.Palette)
		quantizer.Quantize(clonePaletted, canvasBounds, im, image.Point{})
		clonePaletteds[i] = clonePaletted
		disposal := byte(gif.DisposalNone)

		if o.Transparent {
			disposal = byte(gif.DisposalBackground)
		}

		disposals = append(disposals, disposal)
	}

	if reverse && o.Stitch != Shuffle {
		ReverseSlice(clonePaletteds)
		ReverseSlice(sourceDelays)
	}

	clonePaletteds = clonePaletteds[o.TrimStart:]
	clonePaletteds = clonePaletteds[:len(clonePaletteds)-o.TrimEnd]
	sourceDelays = sourceDelays[o.TrimStart:]
	sourceDelays = sourceDelays[:len(sourceDelays)-o.TrimEnd]
	cloneDisposals := disposals[o.TrimStart:]
	cloneDisposals = cloneDisposals[:len(clonePaletteds)-o.TrimEnd]

	if window != 0 {
		clonePaletteds = clonePaletteds[:window]
		sourceDelays = sourceDelays[:window]
		cloneDisposals = cloneDisposals[:window]
	}

	clonePalettedsLen := len(clonePaletteds)

	if o.CutInterval != 0 {
		var reducedPaletteds []*image.Paletted
		var reducedDelays []int
		var reducedDisposals []byte

		for i := 0; i < clonePalettedsLen; i++ {
			if (1+i)%o.CutInterval != 0 {
				reducedPaletteds = append(reducedPaletteds, clonePaletteds[i])
				reducedDelays = append(reducedDelays, sourceDelays[i])
				reducedDisposals = append(reducedDisposals, cloneDisposals[i])
			}
		}

		clonePaletteds = reducedPaletteds
		clonePalettedsLen = len(reducedPaletteds)
		sourceDelays = reducedDelays
		cloneDisposals = reducedDisposals
	}

	var butteryPalettedsLen int

	switch o.Stitch {
	case Mirror:
		butteryPalettedsLen = 2*clonePalettedsLen - 1
	case FlipH:
		butteryPalettedsLen = 2 * clonePalettedsLen
	case FlipV:
		butteryPalettedsLen = 2 * clonePalettedsLen
	default:
		butteryPalettedsLen = clonePalettedsLen
	}

	butteryPaletteds := make([]*image.Paletted, butteryPalettedsLen)
	butteryDelays := make([]int, butteryPalettedsLen)
	butteryDelaysLen := butteryPalettedsLen
	butteryDisposals := make([]byte, butteryPalettedsLen)
	var r int

	for i := 0; i < butteryPalettedsLen; i++ {
		paletted := clonePaletteds[r]

		if (o.Stitch == FlipH || o.Stitch == FlipV) && i > clonePalettedsLen-1 {
			flipPaletted := image.NewPaletted(canvasBounds, nil)

			var flippedRGBA *image.RGBA

			if o.Stitch == FlipH {
				flippedRGBA = transform.FlipH(paletted)
			} else {
				flippedRGBA = transform.FlipV(paletted)
			}

			quantizer.Quantize(flipPaletted, canvasBounds, flippedRGBA, image.Point{})
			paletted = flipPaletted
		}

		butteryPaletteds[i] = paletted
		sourceDelay := sourceDelays[r]
		butteryDelays[i] = int(math.Max(2.0, scaleDelay*float64(sourceDelay)))
		butteryDisposals[i] = cloneDisposals[r]

		switch {
		case o.Stitch == Mirror && i >= clonePalettedsLen-1:
			r--
		case (o.Stitch == FlipH || o.Stitch == FlipV) && i == clonePalettedsLen-1:
			r = 0
		default:
			r++
		}
	}

	if o.Stitch == Shuffle {
		ShuffleSlice(butteryPaletteds)
		ShuffleSlice(butteryDelays)
	} else {
		shiftedPaletteds := make([]*image.Paletted, butteryPalettedsLen)
		shiftedDelays := make([]int, butteryDelaysLen)
		shiftedDisposals := make([]byte, butteryDelaysLen)

		for i := range butteryPaletteds {
			r = (i + o.Shift) % butteryPalettedsLen

			if r < 0 {
				r += butteryPalettedsLen
			}

			shiftedPaletteds[i] = butteryPaletteds[r]
			shiftedDelays[i] = butteryDelays[r]
			shiftedDisposals[i] = butteryDisposals[r]
		}

		butteryPaletteds = shiftedPaletteds
		butteryDelays = shiftedDelays
		butteryDisposals = shiftedDisposals
	}

	butteryGif := gif.GIF{
		LoopCount:       o.LoopCount,
		BackgroundIndex: sourceGif.BackgroundIndex,
		Config:          sourceGif.Config,
		Image:           butteryPaletteds,
		Delay:           butteryDelays,
		Disposal:        butteryDisposals,
	}

	butteryFile, err := os.Create(destPth)
	if err != nil {
		return err
	}

	return gif.EncodeAll(butteryFile, &butteryGif)
}
