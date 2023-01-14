package utils

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestInsertSorted(t *testing.T) {
	f := func(x, y int) int {
		if x < y {
			return -1
		} else if x == y {
			return 0
		} else {
			return 1
		}
	}

	func() {
		s := []int{1, 3}
		x := 2

		got, _ := InsertSorted(s, x, f)
		want := []int{1, 2, 3}

		if slices.Compare(got, want) != 0 {
			t.Errorf("InsertSorted = %v, want %v", got, want)
		}
	}()

	func() {
		s := []int{1, 2}
		x := 3

		got, _ := InsertSorted(s, x, f)
		want := []int{1, 2, 3}

		if slices.Compare(got, want) != 0 {
			t.Errorf("InsertSorted = %v, want %v", got, want)
		}
	}()

	func() {
		s := []int{}
		x := 1

		got, _ := InsertSorted(s, x, f)
		want := []int{1}

		if slices.Compare(got, want) != 0 {
			t.Errorf("InsertSorted = %v, want %v", got, want)
		}
	}()
}
