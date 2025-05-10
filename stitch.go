//go:generate stringer -type=Stitch

// Package buttery provides primitives for manipulating GIF animations.
package buttery

import (
	"fmt"
)

// Stitch models a loop continuity strategy.
type Stitch int

const (
	// None ends the incoming sequence as-is.
	None Stitch = iota

	// Mirror follows the end of the incoming sequence by replaying the sequence backwards.
	Mirror

	// FlipH follows the end of the incoming sequence by replaying the sequence reflected horizontally.
	FlipH

	// FlipV follows the end of the incoming sequence by replaying the sequence reflected vertically.
	FlipV

	// Shuffle randomizes the incoming sequence.
	Shuffle

	// PanH shifts the canvas horizontally
	PanH

	// PanV shifts the canvas vertically
	PanV
)

// ParseStitch generates a Stitch from a string value.
func ParseStitch(s string) (*Stitch, bool) {
	//
	// /!\ Manually update upper bound for each new enum value /!\
	//
	for i := None; i <= PanV; i++ {
		if s == i.String() {
			return &i, true
		}
	}

	return nil, false
}

// Validate rejects out of bound values.
func (o Stitch) Validate() error {
	//
	// /!\ Manually update upper bound for each new enum value /!\
	//
	if o < None || o > PanV {
		return fmt.Errorf("invalid stitch value: %d", o)
	}

	return nil
}
