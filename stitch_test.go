package buttery_test

import (
	"github.com/mcandre/buttery"

	"testing"
)

func TestStitchMarshaling(t *testing.T) {
	stitchString := "Mirror"
	stitch, ok := buttery.ParseStitch(stitchString)

	if !ok {
		t.Errorf("error parsing stitch string %v", stitchString)
	}

	stitchString2 := stitch.String()

	if stitchString2 != stitchString {
		t.Errorf("expected symmetric marshaling for stitch %v", stitch)
	}
}
