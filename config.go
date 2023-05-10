package buttery

import (
	"github.com/andybons/gogif"
	"github.com/disintegration/imaging"

	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"os"
)

// Config models a set of animation editing manipulations.
type Config struct {
	// TrimEdges removes frames from the start and end of the incoming sequence (Default zero).
	TrimEdges int

	// TrimStart removes frames from the start of the incoming sequence (Default zero).
	TrimStart int

	// TrimEnd removes fromes frames from the end of the incoming sequence (Default zero)
	TrimEnd int

	// Window truncates frames from the incoming sequence (Default zero).
	//
	// Zero indicates no window truncation.
	Window int

	// Shift moves the start of the sequence leftward (Default zero).
	Shift int

	// Stitch denotes a loop continuity transition (Default Mirror).
	Stitch Stitch

	// Speed applies a factor to each frame delay (Default 1.0).
	//
	// Zero speed behaves as a Window of a single frame.
	//
	// Negative speed reverses the incoming sequence.
	Speed float64
}

// NewConfig generates a default Config.
func NewConfig() Config {
	return Config{
		Stitch: Mirror,
		Speed:  1.0,
	}
}

// Validate checks for basic Config integrity.
func (o Config) Validate() error {
	if o.TrimEdges < 0 {
		return errors.New("trim edges cannot be negative")
	}

	if o.TrimStart < 0 {
		return errors.New("trim start cannot be negative")
	}

	if o.TrimEnd < 0 {
		return errors.New("trim end cannot be negative")
	}

	if o.Window < 0 {
		return errors.New("window cannot be negative")
	}

	return o.Stitch.Validate()
}

// Edit applies the configured GIF manipulations.
func (o Config) Edit(destPth string, sourceGif *gif.GIF) error {
	sourcePaletteds := sourceGif.Image
	sourcePalettedsLen := len(sourcePaletteds)

	if o.TrimStart+o.TrimEnd >= sourcePalettedsLen {
		return errors.New("minimum 1 output frame")
	}

	var reverse bool
	speed := o.Speed

	if speed < 0 {
		reverse = true
		speed *= -1.0
	}

	window := o.Window

	if speed == 0.0 {
		window = 1
	}

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
	draw.DrawMask(canvasImage, canvasBounds, &image.Uniform{sourcePaletteds[0].Palette.Convert(color.Black)}, image.Point{}, nil, image.Pt(0, 0), draw.Src)

	for i, sourcePaletted := range sourcePaletteds {
		draw.Draw(canvasImage, canvasBounds, sourcePaletted, image.Point{}, draw.Over)
		clonePaletted := image.NewPaletted(canvasBounds, nil)
		quantizer.Quantize(clonePaletted, canvasBounds, canvasImage, image.Point{})
		clonePaletteds[i] = clonePaletted
	}

	if reverse {
		ReverseSlice(clonePaletteds)
		ReverseSlice(sourceDelays)
	}

	clonePaletteds = clonePaletteds[o.TrimStart:]
	clonePaletteds = clonePaletteds[:len(clonePaletteds)-o.TrimEnd]
	sourceDelays = sourceDelays[o.TrimStart:]
	sourceDelays = sourceDelays[:len(sourceDelays)-o.TrimEnd]

	if window != 0 {
		clonePaletteds = clonePaletteds[:window]
		sourceDelays = sourceDelays[:window]
	}

	clonePalettedsLen := len(clonePaletteds)
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
	var r int

	for i := 0; i < butteryPalettedsLen; i++ {
		paletted := clonePaletteds[r]

		if (o.Stitch == FlipH || o.Stitch == FlipV) && i > clonePalettedsLen-1 {
			flipPaletted := image.NewPaletted(canvasBounds, nil)

			var flippedNRGBA *image.NRGBA

			if o.Stitch == FlipH {
				flippedNRGBA = imaging.FlipH(paletted)
			} else {
				flippedNRGBA = imaging.FlipV(paletted)
			}

			quantizer.Quantize(flipPaletted, canvasBounds, flippedNRGBA, image.Point{})
			paletted = flipPaletted
		}

		butteryPaletteds[i] = paletted
		sourceDelay := sourceDelays[r]
		butteryDelays[i] = int(math.Max(2.0, float64(sourceDelay)/speed))

		if o.Stitch == Mirror && i >= clonePalettedsLen-1 {
			r--
		} else if (o.Stitch == FlipH || o.Stitch == FlipV) && i == clonePalettedsLen-1 {
			r = 0
		} else {
			r++
		}
	}

	var shiftedPaletteds = make([]*image.Paletted, butteryPalettedsLen)
	var shiftedDelays = make([]int, butteryDelaysLen)

	for i := range butteryPaletteds {
		r = (i + o.Shift) % butteryPalettedsLen

		if r < 0 {
			r += butteryPalettedsLen
		}

		shiftedPaletteds[i] = butteryPaletteds[r]
		shiftedDelays[i] = butteryDelays[r]
	}

	butteryGif := gif.GIF{
		LoopCount:       0,
		BackgroundIndex: sourceGif.BackgroundIndex,
		Config:          sourceGif.Config,
		Image:           shiftedPaletteds,
		Delay:           shiftedDelays,
		Disposal:        nil,
	}

	butteryFile, err := os.Create(destPth)

	if err != nil {
		return err
	}

	return gif.EncodeAll(butteryFile, &butteryGif)
}
