package main

import (
	"github.com/andybons/gogif"

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
var flagStart = flag.Int("trimStart", 0, "the number of frames to remove from the loop start")
var flagEnd = flag.Int("trimEnd", 0, "the number of frames to remove from the loop end")

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

func main() {
	flag.Parse()

	sourcePth := *flagIn

	if sourcePth == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	trimStart := *flagStart

	if trimStart < 0 {
		panic("trim start cannot be negative")
		os.Exit(1)
	}

	trimEnd := *flagEnd

	if trimEnd < 0 {
		panic("trim end cannot be negative")
		os.Exit(1)
	}

	sourceFile, err := os.Open(sourcePth)

	if err != nil {
		panic(err)
	}

	sourceGif, err := gif.DecodeAll(sourceFile)

	if err != nil {
		panic(err)
	}

	sourcePaletteds := sourceGif.Image

	if trimStart+trimEnd >= len(sourcePaletteds) {
		panic("minimum 1 output frame")
		os.Exit(1)
	}

	sourceDelays := sourceGif.Delay
	sourcePalettedsLen := len(sourcePaletteds)
	sourceWidth, sourceHeight := getDimensions(sourcePaletteds)
	canvasImage := image.NewRGBA(image.Rect(0, 0, sourceWidth, sourceHeight))
	canvasBounds := canvasImage.Bounds()
	paletteSize := getPaletteSize(sourcePaletteds)
	draw.Draw(canvasImage, canvasBounds, sourcePaletteds[0], image.ZP, draw.Src)
	clonePaletteds := make([]*image.Paletted, sourcePalettedsLen)

	for i, sourcePaletted := range sourcePaletteds {
		draw.Draw(canvasImage, canvasBounds, sourcePaletted, image.ZP, draw.Over)
		clonePaletted := image.NewPaletted(canvasBounds, nil)
		quantizer := gogif.MedianCutQuantizer{NumColor: paletteSize}
		quantizer.Quantize(clonePaletted, canvasBounds, canvasImage, image.ZP)
		clonePaletteds[i] = clonePaletted
	}

	clonePaletteds = clonePaletteds[trimStart:len(sourcePaletteds)-trimEnd]
	clonePalettedsLen := len(clonePaletteds)
	sourceDelays = sourceDelays[trimStart:len(sourceDelays)-trimEnd]

	butteryPalettedsLen := 2*(clonePalettedsLen)-1
	butteryPaletteds := make([]*image.Paletted, butteryPalettedsLen)
	butteryDelays := make([]int, butteryPalettedsLen)

	m := butteryPalettedsLen/2

	for i := 0; i <= m; i++ {
		butteryPaletteds[i] = clonePaletteds[i]
		butteryDelays[i] = sourceDelays[i]
	}

	for i := m + 1; i < butteryPalettedsLen; i++ {
		r := butteryPalettedsLen - i - 1
		butteryPaletteds[i] = butteryPaletteds[r]
		butteryDelays[i] = butteryDelays[r]
	}

	butteryGif := gif.GIF{
		LoopCount: 0,
		BackgroundIndex: sourceGif.BackgroundIndex,
		Config: sourceGif.Config,
		Image: butteryPaletteds,
		Delay: butteryDelays,
		Disposal: nil,
	}

	sourceBasename := strings.TrimSuffix(sourcePth, filepath.Ext(sourcePth))
	butteryPth := fmt.Sprintf("%v.buttery.gif", sourceBasename)
	butteryFile, err := os.Create(butteryPth)

	if err != nil {
		panic(err)
	}

	if err2 := gif.EncodeAll(butteryFile, &butteryGif); err2 != nil {
		panic(err2)
	}
}
