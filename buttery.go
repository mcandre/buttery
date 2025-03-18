// Package buttery provides primitives for manipulating GIF animations.
package buttery

import (
	"image"
	"math/rand"
	"reflect"
)

// ReverseSlice performs an in-place swap in reverse order.
func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)

	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// ShuffleSlice randomizes the order of a slice.
func ShuffleSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	rand.Shuffle(size, swap)
}

// GetDimensions reports the horizontal and vertical bounds of a GIF.
func GetDimensions(paletteds []*image.Paletted) (int, int) {
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

// GetPaletteSize queries the size of a GIF's color space.
func GetPaletteSize(paletteds []*image.Paletted) int {
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
