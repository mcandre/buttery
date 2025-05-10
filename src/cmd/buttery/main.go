// Package main implements a GIF editor application.
package main

import (
	"github.com/mcandre/buttery"

	"flag"
	"fmt"
	"image/gif"
	"os"
	"path/filepath"
	"strings"
)

var flagCheck = flag.Bool("check", false, "validate basic GIF format file integrity")
var flagGetFrames = flag.Bool("getFrames", false, "query total input GIF frame count")
var flagTransparent = flag.Bool("transparent", false, "preserve clear GIFs")
var flagTrimEdges = flag.Int("trimEdges", 0, "drop frames from both ends of the input GIF")
var flagTrimStart = flag.Int("trimStart", 0, "drop frames from start of the input GIF")
var flagTrimEnd = flag.Int("trimEnd", 0, "drop frames from end of the input GIF")
var flagCutInterval = flag.Int("cutInterval", 0, "drop every nth frame of the input GIF")
var flagWindow = flag.Int("window", 0, "set fixed sequence length")
var flagStitch = flag.String("stitch", "Mirror", "stitching strategy (None/Mirror/FlipH/FlipV/Shuffle/PanH/PanV)")
var flagShift = flag.Int("shift", 0, "rotate sequence left")
var flagScaleDelay = flag.Float64("scaleDelay", 1.0, "multiply each frame delay by a factor")
var flagPanVelocity = flag.Float64("panVelocity", 1, "how many pixels to pan per frame")
var flagLoopCount = flag.Int("loopCount", 0, "how many times to play animation (-1: Once, 0: Infinite, N: N+1 iterations)")
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
	config.Transparent = *flagTransparent
	config.TrimStart = *flagTrimStart + trimEdges
	config.TrimEnd = *flagTrimEnd + trimEdges
	config.CutInterval = *flagCutInterval
	config.Window = *flagWindow
	config.Shift = *flagShift
	config.Stitch = *stitchP
	config.ScaleDelay = *flagScaleDelay
	config.PanVelocity = *flagPanVelocity
	config.LoopCount = *flagLoopCount

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

	if getFrames {
		fmt.Println(len(sourcePaletteds))
		os.Exit(0)
	}

	sourceBasename := strings.TrimSuffix(sourcePth, filepath.Ext(sourcePth))
	destPth := fmt.Sprintf("%v.buttery.gif", sourceBasename)

	if err := config.Edit(destPth, sourceGif); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
