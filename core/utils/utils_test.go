package utils

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestInsertSorted(t *testing.T) {
	s := []int{1, 3}
	x := 2
	f := func(x, y int) int {
		if x < y {
			return -1
		} else if x == y {
			return 0
		} else {
			return 1
		}
	}

	InsertSorted(s, x, f)
	want := []int{1, 2, 3}

	if slices.Compare(s, want) != 0 {
		t.Errorf("InsertSorted = %v, want %v", s, want)
	}
}
