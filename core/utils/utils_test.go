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

func TestXorSlice(t *testing.T) {
	f := func(x, y int) int {
		if x < y {
			return -1
		} else if x == y {
			return 0
		} else {
			return 1
		}
	}

	// func() {
	// 	// s1 has many elements
	// 	s1 := []int{1, 2, 3, 4}
	// 	s2 := []int{2, 3}

	// 	got := XorSlice(s1, s2, f)
	// 	want := []int{1, 4}

	// 	if slices.Compare(got, want) != 0 {
	// 		t.Errorf("s1 has many elements failed: InsertSorted = %v, want %v", got, want)
	// 	}
	// }()

	func() {
		// s2 has many elements
		s1 := []int{2, 4}
		s2 := []int{1, 2, 3, 4, 5}

		got := XorSlice(s1, s2, f)
		want := []int{1, 3, 5}

		if slices.Compare(got, want) != 0 {
			t.Errorf("s2 has many elements failed: InsertSorted = %v, want %v", got, want)
		}
	}()

	// func() {
	// 	// both have many elements
	// 	s1 := []int{1, 2, 2, 4, 4, 6}
	// 	s2 := []int{1, 1, 2, 3, 5}

	// 	got := XorSlice(s1, s2, f)
	// 	want := []int{3, 4, 5, 6}

	// 	if slices.Compare(got, want) != 0 {
	// 		t.Errorf("both have many elements failed: InsertSorted = %v, want %v", got, want)
	// 	}
	// }()

	// func() {
	// 	// same slice
	// 	s1 := []int{1, 2, 2, 3}
	// 	s2 := []int{1, 1, 2, 3}

	// 	got := XorSlice(s1, s2, f)
	// 	want := []int{}

	// 	if slices.Compare(got, want) != 0 {
	// 		t.Errorf("same slice failed: InsertSorted = %v, want %v", got, want)
	// 	}
	// }()
}
