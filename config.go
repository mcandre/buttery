package buttery

import (
	"errors"
)

// Config models a set of animation editing instructions.
type Config struct {
	// Reverse plays the incoming sequence backwards (Default false).
	Reverse bool

	// TrimEdges removes frames from the start and end of the incoming sequence (Default zero).
	TrimEdges int

	// TrimStart removes frames from the start of the incoming sequence (Default zero).
	TrimStart int

	// TrimEnd removes fromes frames from the end of the incoming sequence (Default zero)
	TrimEnd int

	// Window truncates frames from the incoming sequence (Default zero).
	//
	// Zero indicates no window truncation.
	Window int

	// Shift moves the start of the sequence leftward (Default zero).
	Shift int

	// Stitch denotes a loop continuity transition (Default Mirror).
	Stitch Stitch

	//
	Speed float64
}

// NewConfig generates a default Config.
func NewConfig() Config {
	return Config{
		Stitch: Mirror,
		Speed:  1.0,
	}
}

// Validate checks for basic Config integrity.
func (o Config) Validate() error {
	if o.TrimEdges < 0 {
		return errors.New("trim edges cannot be negative")
	}

	if o.TrimStart < 0 {
		return errors.New("trim start cannot be negative")
	}

	if o.TrimEnd < 0 {
		return errors.New("trim end cannot be negative")
	}

	if o.Window < 0 {
		return errors.New("window cannot be negative")
	}

	if o.Speed <= 0 {
		return errors.New("speed must be positive")
	}

	return nil
}
