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
var flagEdges = flag.Int("trimEdges", 0, "drop frames from both ends of the input GIF")
var flagStart = flag.Int("trimStart", 0, "drop frames from start of the input GIF")
var flagEnd = flag.Int("trimEnd", 0, "drop frames from end of the input GIF")
var flagWindow = flag.Int("window", -1, "set fixed sequence length")
var flagStitch = flag.String("stitch", "Mirror", "stitching strategy (None/Mirror/FlipH/FlipV)")
var flagReverse = flag.Bool("reverse", false, "reverse original sequence")
var flagShift = flag.Int("shift", 0, "rotate sequence left")
var flagSpeed = flag.Float64("speed", 1.0, "speed factor (highly sensitive)")
var flagVersion = flag.Bool("version", false, "show version information")
var flagHelp = flag.Bool("help", false, "show usage information")

func Usage() {
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
		Usage()
		os.Exit(0)
	}

	if *flagVersion {
		fmt.Println(buttery.Version)
		os.Exit(0)
	}

	check := *flagCheck
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
	stitchString := *flagStitch
	stitchP, ok := buttery.ParseStitch(stitchString)

	if !ok {
		Usage()
		os.Exit(1)
	}

	stitch := *stitchP

	if *flagSpeed <= 0.0 {
		fmt.Fprintln(os.Stderr, "speed must be positive")
		os.Exit(1)
	}

	speed := *flagSpeed
	rest := flag.Args()

	if len(rest) != 1 {
		Usage()
		os.Exit(1)
	}

	sourcePth := rest[0]

	if sourcePth == "" {
		Usage()
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

	if trimStart+trimEnd >= sourcePalettedsLen {
		fmt.Fprintln(os.Stderr, "minimum 1 output frame")
		os.Exit(1)
	}

	if window > sourcePalettedsLen-trimStart-trimEnd {
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

	if reverse {
		buttery.ReverseSlice(clonePaletteds)
		buttery.ReverseSlice(sourceDelays)
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

	switch stitch {
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

		if (stitch == buttery.FlipH || stitch == buttery.FlipV) && i > clonePalettedsLen-1 {
			flipPaletted := image.NewPaletted(canvasBounds, nil)

			var flippedNRGBA *image.NRGBA

			if stitch == buttery.FlipH {
				flippedNRGBA = imaging.FlipH(paletted)
			} else {
				flippedNRGBA = imaging.FlipV(paletted)
			}

			quantizer.Quantize(flipPaletted, canvasBounds, flippedNRGBA, image.ZP)
			paletted = flipPaletted
		}

		butteryPaletteds[i] = paletted
		sourceDelay := sourceDelays[r]
		butteryDelays[i] = int(math.Max(2.0, float64(sourceDelay)/speed))

		if stitch == buttery.Mirror && i >= clonePalettedsLen-1 {
			r--
		} else if (stitch == buttery.FlipH || stitch == buttery.FlipV) && i == clonePalettedsLen-1 {
			r = 0
		} else {
			r++
		}
	}

	var shiftedPaletteds = make([]*image.Paletted, butteryPalettedsLen)
	var shiftedDelays = make([]int, butteryDelaysLen)

	for i := range butteryPaletteds {
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
