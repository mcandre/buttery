package main

import (
	"github.com/andybons/gogif"
	"github.com/disintegration/imaging"
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

var flagCheck = flag.Bool("check", false, "validate basic GIF format file integrity")
var flagGetFrames = flag.Bool("getFrames", false, "query total input GIF frame count")
var flagTrimEdges = flag.Int("trimEdges", 0, "drop frames from both ends of the input GIF")
var flagTrimStart = flag.Int("trimStart", 0, "drop frames from start of the input GIF")
var flagTrimEnd = flag.Int("trimEnd", 0, "drop frames from end of the input GIF")
var flagWindow = flag.Int("window", 0, "set fixed sequence length")
var flagStitch = flag.String("stitch", "Mirror", "stitching strategy (None/Mirror/FlipH/FlipV)")
var flagReverse = flag.Bool("reverse", false, "reverse original sequence")
var flagShift = flag.Int("shift", 0, "rotate sequence left")
var flagSpeed = flag.Float64("speed", 1.0, "speed factor (highly sensitive)")
var flagVersion = flag.Bool("version", false, "show version information")
var flagHelp = flag.Bool("help", false, "show usage information")

func usage() {
	program, err := os.Executable()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Usage: %v [OPTION] <input.gif>\n", program)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if *flagHelp {
		usage()
		os.Exit(0)
	}

	if *flagVersion {
		fmt.Println(buttery.Version)
		os.Exit(0)
	}

	rest := flag.Args()

	if len(rest) != 1 {
		usage()
		os.Exit(1)
	}

	sourcePth := rest[0]

	if sourcePth == "" {
		usage()
		os.Exit(1)
	}

	check := *flagCheck
	getFrames := *flagGetFrames
	trimEdges := *flagTrimEdges

	if trimEdges < 0 {
		fmt.Fprintln(os.Stderr, "trim edges cannot be negative")
		os.Exit(1)
	}

	stitchString := *flagStitch
	stitchP, ok := buttery.ParseStitch(stitchString)

	if !ok {
		usage()
		os.Exit(1)
	}

	config := buttery.NewConfig()
	config.Reverse = *flagReverse
	config.TrimStart = *flagTrimStart + trimEdges
	config.TrimEnd = *flagTrimEnd + trimEdges
	config.Window = *flagWindow
	config.Shift = *flagShift
	config.Stitch = *stitchP
	config.Speed = *flagSpeed

	if err := config.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
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

	if check {
		os.Exit(0)
	}

	sourcePaletteds := sourceGif.Image
	sourcePalettedsLen := len(sourcePaletteds)

	if getFrames {
		fmt.Println(sourcePalettedsLen)
		os.Exit(0)
	}

	if config.TrimStart+config.TrimEnd >= sourcePalettedsLen {
		fmt.Fprintln(os.Stderr, "minimum 1 output frame")
		os.Exit(1)
	}

	if config.Window > sourcePalettedsLen-config.TrimStart-config.TrimEnd {
		fmt.Fprintln(os.Stderr, "window longer than subsequence")
		os.Exit(1)
	}

	sourceDelays := sourceGif.Delay
	sourceWidth, sourceHeight := buttery.GetDimensions(sourcePaletteds)
	canvasImage := image.NewRGBA(image.Rect(0, 0, sourceWidth, sourceHeight))
	canvasBounds := canvasImage.Bounds()
	paletteSize := buttery.GetPaletteSize(sourcePaletteds)
	clonePaletteds := make([]*image.Paletted, sourcePalettedsLen)
	quantizer := gogif.MedianCutQuantizer{NumColor: paletteSize}
	draw.DrawMask(canvasImage, canvasBounds, &image.Uniform{sourcePaletteds[0].Palette.Convert(color.Black)}, image.ZP, nil, image.Pt(0, 0), draw.Src)

	for i, sourcePaletted := range sourcePaletteds {
		draw.Draw(canvasImage, canvasBounds, sourcePaletted, image.ZP, draw.Over)
		clonePaletted := image.NewPaletted(canvasBounds, nil)
		quantizer.Quantize(clonePaletted, canvasBounds, canvasImage, image.ZP)
		clonePaletteds[i] = clonePaletted
	}

	if config.Reverse {
		buttery.ReverseSlice(clonePaletteds)
		buttery.ReverseSlice(sourceDelays)
	}

	clonePaletteds = clonePaletteds[config.TrimStart:]
	clonePaletteds = clonePaletteds[:len(clonePaletteds)-config.TrimEnd]
	sourceDelays = sourceDelays[config.TrimStart:]
	sourceDelays = sourceDelays[:len(sourceDelays)-config.TrimEnd]

	if config.Window != 0 {
		clonePaletteds = clonePaletteds[:config.Window]
		sourceDelays = sourceDelays[:config.Window]
	}

	clonePalettedsLen := len(clonePaletteds)
	var butteryPalettedsLen int

	switch config.Stitch {
	case buttery.Mirror:
		butteryPalettedsLen = 2*clonePalettedsLen - 1
	case buttery.FlipH:
		butteryPalettedsLen = 2 * clonePalettedsLen
	case buttery.FlipV:
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

		if (config.Stitch == buttery.FlipH || config.Stitch == buttery.FlipV) && i > clonePalettedsLen-1 {
			flipPaletted := image.NewPaletted(canvasBounds, nil)

			var flippedNRGBA *image.NRGBA

			if config.Stitch == buttery.FlipH {
				flippedNRGBA = imaging.FlipH(paletted)
			} else {
				flippedNRGBA = imaging.FlipV(paletted)
			}

			quantizer.Quantize(flipPaletted, canvasBounds, flippedNRGBA, image.ZP)
			paletted = flipPaletted
		}

		butteryPaletteds[i] = paletted
		sourceDelay := sourceDelays[r]
		butteryDelays[i] = int(math.Max(2.0, float64(sourceDelay)/config.Speed))

		if config.Stitch == buttery.Mirror && i >= clonePalettedsLen-1 {
			r--
		} else if (config.Stitch == buttery.FlipH || config.Stitch == buttery.FlipV) && i == clonePalettedsLen-1 {
			r = 0
		} else {
			r++
		}
	}

	var shiftedPaletteds = make([]*image.Paletted, butteryPalettedsLen)
	var shiftedDelays = make([]int, butteryDelaysLen)

	for i := range butteryPaletteds {
		r = (i + config.Shift) % butteryPalettedsLen

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
