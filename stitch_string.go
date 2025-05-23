// Code generated by "stringer -type Stitch"; DO NOT EDIT.

package buttery

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[None-0]
	_ = x[Mirror-1]
	_ = x[FlipH-2]
	_ = x[FlipV-3]
	_ = x[Shuffle-4]
	_ = x[PanH-5]
	_ = x[PanV-6]
}

const _Stitch_name = "NoneMirrorFlipHFlipVShufflePanHPanV"

var _Stitch_index = [...]uint8{0, 4, 10, 15, 20, 27, 31, 35}

func (i Stitch) String() string {
	if i < 0 || i >= Stitch(len(_Stitch_index)-1) {
		return "Stitch(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Stitch_name[_Stitch_index[i]:_Stitch_index[i+1]]
}
