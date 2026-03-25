package buttery

import (
	"image"
)

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
