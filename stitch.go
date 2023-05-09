//go:generate stringer -type=Stitch
package buttery

// Stitch models a loop continuity strategy.
type Stitch int

const (
	// Mirror follows the end of the incoming sequence by replaying the sequence backwards.
	Mirror Stitch = iota

	// Flip follows the end of the incoming sequence by replaying the sequence reflected horizontally.
	Flip

	// None ends the incoming sequence as-is.
	None
)

// ParseStitch generates a Stitch from a string value.
func ParseStitch(s string) (*Stitch, bool) {
	for i := Mirror; i <= None; i++ {
		if s == i.String() {
			return &i, true
		}
	}

	return nil, false
}
