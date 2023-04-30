package main

import (
	"reflect"

	"github.com/andybons/gogif"
	"github.com/mcandre/buttery"

	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"path/filepath"
	"strings"
)

var flagIn = flag.String("in", "", "path to a .gif source file (required)")
var flagGetFrames = flag.Bool("getFrames", false, "query total input GIF frame count")
var flagStart = flag.Int("trimStart", 0, "drop frames from start of the input GIF")
var flagEnd = flag.Int("trimEnd", 0, "drop frames from end of the input GIF")
var flagMirror = flag.Bool("mirror", true, "Toggle frame sequence mirroring")
var flagReverse = flag.Bool("reverse", false, "Reverse original sequence")
var flagVersion = flag.Bool("version", false, "Show version information")
var flagHelp = flag.Bool("help", false, "Show usage information")

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

	reverse := *flagReverse
	mirror := *flagMirror

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

	sourceDelays := sourceGif.Delay
	sourceWidth, sourceHeight := getDimensions(sourcePaletteds)
	canvasImage := image.NewRGBA(image.Rect(0, 0, sourceWidth, sourceHeight))
	canvasBounds := canvasImage.Bounds()
	paletteSize := getPaletteSize(sourcePaletteds)
	clonePaletteds := make([]*image.Paletted, sourcePalettedsLen)

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
	clonePalettedsLen := len(clonePaletteds)

	sourceDelays = sourceDelays[trimStart:]
	var butteryPalettedsLen int

	if mirror {
		butteryPalettedsLen = 2*clonePalettedsLen - 1
	} else {
		butteryPalettedsLen = clonePalettedsLen
	}

	butteryPaletteds := make([]*image.Paletted, butteryPalettedsLen)
	butteryDelays := make([]int, butteryPalettedsLen)
	var r int

	for i := 0; i < butteryPalettedsLen; i++ {
		butteryPaletteds[i] = clonePaletteds[r]
		butteryDelays[i] = sourceDelays[r]

		if !mirror || i < clonePalettedsLen-1 {
			r++
		} else {
			r--
		}
	}

	butteryGif := gif.GIF{
		LoopCount:       0,
		BackgroundIndex: sourceGif.BackgroundIndex,
		Config:          sourceGif.Config,
		Image:           butteryPaletteds,
		Delay:           butteryDelays,
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
