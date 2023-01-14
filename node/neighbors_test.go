package main

import (
	"local-fog/core/types"
	"testing"
)

func TestInsetNeighbor(t *testing.T) {
	comp := func(s1, s2 []*types.NodeInfoWrapper) bool {
		if len(s1) != len(s2) {
			return false
		}

		for i := range s1 {
			if types.CompareNode(s1[i], s2[i]) != 0 {
				return false
			}
		}

		return true
	}

	func() {
		n1 := &types.NodeInfoWrapper{Id: 1}
		n3 := &types.NodeInfoWrapper{Id: 3}
		s := []*types.NodeInfoWrapper{n1, n3}
		x := &types.NodeInfoWrapper{Id: 2}

		UpdateNeighbors(s)

		InsertNeighbor(x)
		want := []*types.NodeInfoWrapper{n1, x, n3}
		got := Neighbors

		if !comp(want, got) {
			t.Errorf("InsertNeighbor = %v, want %v", got, want)
		}
	}()

	func() {
		n2 := &types.NodeInfoWrapper{Id: 2}
		n3 := &types.NodeInfoWrapper{Id: 3}
		s := []*types.NodeInfoWrapper{n2, n3}
		x := &types.NodeInfoWrapper{Id: 1}

		UpdateNeighbors(s)

		InsertNeighbor(x)
		want := []*types.NodeInfoWrapper{x, n2, n3}
		got := Neighbors

		if !comp(want, got) {
			t.Errorf("InsertNeighbor = %v, want %v", got, want)
		}
	}()

	func() {
		s := []*types.NodeInfoWrapper{}
		x := &types.NodeInfoWrapper{Id: 1}

		UpdateNeighbors(s)

		InsertNeighbor(x)
		want := []*types.NodeInfoWrapper{x}
		got := Neighbors

		if !comp(want, got) {
			t.Errorf("InsertNeighbor = %v, want %v", got, want)
		}
	}()
}
