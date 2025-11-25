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

	// FadeColor denotes a hue for fade transitions (Default opaque black).
	// Alpha channel ignored.
	FadeColor color.RGBA

	// FadeRate denotes the speed of fade transitions (Default 1.0).
	FadeRate float64

	// ScaleDelay multiplies each frame delay by a factor (Default 1.0).
	//
	// The resulting delay is upheld to a lower bound of 2 centisec.
	//
	// A negative scale delay reverses the incoming sequence.
	ScaleDelay float64

	// PanVelocity specifies the number of pixels to shift the canvas per frame (Default: 1.0).
	PanVelocity float64

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

		panVelocity := int(float64(i) * o.PanVelocity)

		if o.Stitch == PanH {
			panPaletted := pan(paletted, panVelocity, 0)
			paletted = panPaletted
		}

		if o.Stitch == PanV {
			panPaletted := pan(paletted, 0, panVelocity)
			paletted = panPaletted
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
		target := o.FadeColor
		targetR, targetG, targetB := float64(target.R), float64(target.G), float64(target.B)
		s := float64(butteryPalettedsLen) - 1.0

		for i := range butteryPaletteds {
			r = signedMod(i+o.Shift, butteryPalettedsLen)
			shiftedPaletteds[i] = butteryPaletteds[r]
			shiftedDelays[i] = butteryDelays[r]
			shiftedDisposals[i] = butteryDisposals[r]

			fadedPaletted := shiftedPaletteds[i]
			fade := float64(s) / float64(butteryPalettedsLen-1)
			palette := fadedPaletted.Palette
			fadedPalette := make(color.Palette, len(palette))

			for j, c := range palette {
				r, g, b, a := c.RGBA()
				rF, gF, bF := float64(r>>8), float64(g>>8), float64(b>>8)
				rF = rF + (targetR-rF)*fade
				rF = max(rF, 0.0)
				rF = min(rF, 255.0)
				gF = gF + (targetG-gF)*fade
				gF = max(gF, 0.0)
				gF = min(gF, 255.0)
				bF = bF + (targetB-bF)*fade
				bF = max(bF, 0.0)
				bF = min(bF, 255.0)
				r2, g2, b2, a2 := uint8(rF), uint8(gF), uint8(bF), uint8(a>>8)
				fadedPalette[j] = color.RGBA{R: r2, G: g2, B: b2, A: a2}
			}

			fadedPaletted.Palette = fadedPalette
			shiftedPaletteds[i] = fadedPaletted

			if i < butteryPalettedsLen/2 {
				s -= o.FadeRate
			} else if i > butteryPalettedsLen/2 {
				s += o.FadeRate
			}

			s = min(s, float64(butteryPalettedsLen)-1.0)
			s = max(s, 0.0)
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

// pan offsets an image by the given horizontal and vertical offsets.
func pan(img *image.Paletted, dx, dy int) *image.Paletted {
	bounds := img.Bounds()
	panned := image.NewPaletted(bounds, img.Palette)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			srcX, srcY := signedMod(x, bounds.Max.X), signedMod(y, bounds.Max.Y)
			dstX, dstY := signedMod(x+dx, bounds.Max.X), signedMod(y+dy, bounds.Max.Y)

			panned.SetColorIndex(dstX, dstY, img.ColorIndexAt(srcX, srcY))
		}
	}

	return panned
}

// signedMod reverses direction for negative n denominators.
//
// Warning: Each programming language may implements subtly distinct modulo algorithms.
// https://en.wikipedia.org/wiki/Modulo
func signedMod(a, n int) int {
	return a - (n * int(math.Floor(float64(a)/float64(n))))
}
