//go:generate stringer -type=Stitch
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
)

// ParseStitch generates a Stitch from a string value.
func ParseStitch(s string) (*Stitch, bool) {
	for i := None; i <= FlipV; i++ {
		if s == i.String() {
			return &i, true
		}
	}

	return nil, false
}

// Validate rejects out of bound values.
func (o Stitch) Validate() error {
	if o < None || o > FlipV {
		return fmt.Errorf("invalid stitch value: %d", o)
	}

	return nil
}
