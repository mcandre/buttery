package main

import (
	"reflect"

	"github.com/andybons/gogif"
	"github.com/mcandre/buttery"

	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"os"
	"path/filepath"
	"strings"
)

var flagIn = flag.String("in", "", "path to a .gif source file (required)")
var flagGetFrames = flag.Bool("getFrames", false, "query total input GIF frame count")
var flagEdges = flag.Int("trimEdges", 0, "drop frames from both ends of the input GIF")
var flagStart = flag.Int("trimStart", 0, "drop frames from start of the input GIF")
var flagEnd = flag.Int("trimEnd", 0, "drop frames from end of the input GIF")
var flagWindow = flag.Int("window", -1, "set fixed sequence length")
var flagMirror = flag.Bool("mirror", true, "toggle frame sequence mirroring")
var flagReverse = flag.Bool("reverse", false, "reverse original sequence")
var flagShift = flag.Int("shift", 0, "rotate sequence left")
var flagSpeed = flag.Float64("speed", 1.0, "speed factor (highly sensitive)")
var flagVersion = flag.Bool("version", false, "show version information")
var flagHelp = flag.Bool("help", false, "show usage information")

func getDimensions(paletteds []*image.Paletted) (int, int) {
	var xMin int
	var xMax int
	var yMin int
	var yMax int

	for _, paletted := range paletteds {
		rect := paletted.Rect
		rectXMin := rect.Min.X
		rectXMax := rect.Max.X
		rectYMin := rect.Min.Y
		rectYMax := rect.Max.Y

		if rectXMin < xMin {
			xMin = rectXMin
		}

		if rectXMax > xMax {
			xMax = rectXMax
		}

		if rectYMin < yMin {
			yMin = rectYMin
		}

		if rectYMax > yMax {
			yMax = rectYMax
		}
	}

	return xMax - xMin, yMax - yMin
}

func getPaletteSize(paletteds []*image.Paletted) int {
	var maxPaletteSize int

	for _, paletted := range paletteds {
		palette := paletted.Palette
		paletteSize := len(palette)

		if paletteSize > maxPaletteSize {
			maxPaletteSize = paletteSize
		}
	}

	return maxPaletteSize
}

// ReverseSlice performs an in-place swap in reverse order.
func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func main() {
	flag.Parse()

	if *flagHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *flagVersion {
		fmt.Println(buttery.Version)
		os.Exit(0)
	}

	sourcePth := *flagIn

	if sourcePth == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	getFrames := *flagGetFrames
	trimEdges := *flagEdges

	if trimEdges < 0 {
		fmt.Fprintln(os.Stderr, "trim edges cannot be negative")
		os.Exit(1)
	}

	trimStart := *flagStart

	if trimStart < 0 {
		fmt.Fprintln(os.Stderr, "trim start cannot be negative")
		os.Exit(1)
	}

	trimEnd := *flagEnd

	if trimEnd < 0 {
		fmt.Fprintln(os.Stderr, "trim end cannot be negative")
		os.Exit(1)
	}

	trimStart += trimEdges
	trimEnd += trimEdges
	window := *flagWindow

	if window != -1 {
		if window < 1 {
			fmt.Fprintln(os.Stderr, "minimum 1 output frame")
			os.Exit(1)
		}
	}

	reverse := *flagReverse
	shift := *flagShift
	mirror := *flagMirror

	if *flagSpeed <= 0.0 {
		fmt.Fprintln(os.Stderr, "speed must be positive")
		os.Exit(1)
	}

	speed := *flagSpeed

	if len(flag.Args()) > 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	sourceFile, err := os.Open(sourcePth)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	sourceGif, err := gif.DecodeAll(sourceFile)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	sourcePaletteds := sourceGif.Image
	sourcePalettedsLen := len(sourcePaletteds)

	if getFrames {
		fmt.Println(sourcePalettedsLen)
		os.Exit(0)
	}

	if trimStart+trimEnd >= sourcePalettedsLen {
		fmt.Fprintln(os.Stderr, "minimum 1 output frame")
		os.Exit(1)
	}

	if window > sourcePalettedsLen-trimStart-trimEnd {
		fmt.Fprintln(os.Stderr, "window longer than subsequence")
		os.Exit(1)
	}

	sourceDelays := sourceGif.Delay
	sourceWidth, sourceHeight := getDimensions(sourcePaletteds)
	canvasImage := image.NewRGBA(image.Rect(0, 0, sourceWidth, sourceHeight))
	canvasBounds := canvasImage.Bounds()
	paletteSize := getPaletteSize(sourcePaletteds)
	clonePaletteds := make([]*image.Paletted, sourcePalettedsLen)
	draw.DrawMask(canvasImage, canvasBounds, &image.Uniform{sourcePaletteds[0].Palette.Convert(color.Black)}, image.ZP, nil, image.Pt(0, 0), draw.Src)

	for i, sourcePaletted := range sourcePaletteds {
		draw.Draw(canvasImage, canvasBounds, sourcePaletted, image.ZP, draw.Over)
		clonePaletted := image.NewPaletted(canvasBounds, nil)
		quantizer := gogif.MedianCutQuantizer{NumColor: paletteSize}
		quantizer.Quantize(clonePaletted, canvasBounds, canvasImage, image.ZP)
		clonePaletteds[i] = clonePaletted
	}

	if reverse {
		ReverseSlice(clonePaletteds)
	}

	clonePaletteds = clonePaletteds[trimStart:]
	clonePaletteds = clonePaletteds[:len(clonePaletteds)-trimEnd]

	if window != -1 {
		clonePaletteds = clonePaletteds[:window]
	}

	clonePalettedsLen := len(clonePaletteds)
	sourceDelays = sourceDelays[trimStart:]
	sourceDelays = sourceDelays[:len(sourceDelays)-trimEnd]

	if window != -1 {
		sourceDelays = sourceDelays[:window]
	}

	var butteryPalettedsLen int

	if mirror {
		butteryPalettedsLen = 2*clonePalettedsLen - 1
	} else {
		butteryPalettedsLen = clonePalettedsLen
	}

	butteryPaletteds := make([]*image.Paletted, butteryPalettedsLen)
	butteryDelays := make([]int, butteryPalettedsLen)
	butteryDelaysLen := butteryPalettedsLen
	var r int

	for i := 0; i < butteryPalettedsLen; i++ {
		butteryPaletteds[i] = clonePaletteds[r]
		sourceDelay := sourceDelays[r]
		butteryDelay := int(math.Max(2.0, float64(sourceDelay) / speed))
		butteryDelays[i] = butteryDelay

		if !mirror || i < clonePalettedsLen-1 {
			r++
		} else {
			r--
		}
	}

	var shiftedPaletteds = make([]*image.Paletted, butteryPalettedsLen)
	var shiftedDelays = make([]int, butteryDelaysLen)

	for i, _ := range butteryPaletteds {
		r = (i + shift) % butteryPalettedsLen

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

	sourceBasename := strings.TrimSuffix(sourcePth, filepath.Ext(sourcePth))
	butteryPth := fmt.Sprintf("%v.buttery.gif", sourceBasename)
	butteryFile, err := os.Create(butteryPth)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err2 := gif.EncodeAll(butteryFile, &butteryGif); err2 != nil {
		fmt.Fprintln(os.Stderr, err2)
		os.Exit(1)
	}
}
