package buttery_test

import (
	"github.com/mcandre/buttery"

	"reflect"
	"testing"
)

func TestReverseSliceSymmetric(t *testing.T) {
	xs := []int{1, 2, 3}
	buttery.ReverseSlice(xs)
	buttery.ReverseSlice(xs)

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(expected, xs) {
		t.Errorf("expected ReverseSlice of %v to be symmetric", xs)
	}
}
