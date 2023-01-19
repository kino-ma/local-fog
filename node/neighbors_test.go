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

func TestChooseMonitorTarget(t *testing.T) {
	func() {
		// self is n2
		n1 := &types.NodeInfoWrapper{Id: 1}
		n2 := &types.NodeInfoWrapper{Id: 2}
		n3 := &types.NodeInfoWrapper{Id: 3}
		s := []*types.NodeInfoWrapper{n1, n2, n3}

		UpdateNeighbors(s)

		self := n2
		want := n1
		got := chooseMonitorTarget(s, self.Id)

		if types.CompareNode(want, got) != 0 {
			t.Errorf("self is n2: chooseMonitorTarget = %v, want %v", got, want)
		}
	}()

	func() {
		// self is n1
		n1 := &types.NodeInfoWrapper{Id: 1}
		n2 := &types.NodeInfoWrapper{Id: 2}
		n3 := &types.NodeInfoWrapper{Id: 3}
		s := []*types.NodeInfoWrapper{n1, n2, n3}

		UpdateNeighbors(s)

		self := n1
		want := n3
		got := chooseMonitorTarget(s, self.Id)

		if types.CompareNode(want, got) != 0 {
			t.Errorf("self is n1: chooseMonitorTarget = %v, want %v", got, want)
		}
	}()

	func() {
		// length is 1
		n1 := &types.NodeInfoWrapper{Id: 1}
		s := []*types.NodeInfoWrapper{n1}

		UpdateNeighbors(s)

		self := n1
		want := (*types.NodeInfoWrapper)(nil)
		got := chooseMonitorTarget(s, self.Id)

		if want != got {
			t.Errorf("length is 1: hooseMonitorTarget = %v, want %v", got, want)
		}
	}()
}

func TestPatchNeighbors(t *testing.T) {
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
		// no common nodes
		n1 := &types.NodeInfoWrapper{Id: 1}
		n2 := &types.NodeInfoWrapper{Id: 2}
		n3 := &types.NodeInfoWrapper{Id: 3}
		n4 := &types.NodeInfoWrapper{Id: 4}
		n5 := &types.NodeInfoWrapper{Id: 1}
		n6 := &types.NodeInfoWrapper{Id: 1}

		s1 := []*types.NodeInfoWrapper{n1, n2, n3}
		s2 := []*types.NodeInfoWrapper{n4, n5, n6}

		got := PatchNeighbors(s1, s2)
		want := []*types.NodeInfoWrapper{n1, n2, n3, n4, n5, n6}

		if !comp(want, got) {
			t.Errorf("no common nodes: chooseMonitorTarget = %v, want %v", got, want)
		}
	}()

	func() {
		// 1 common node
		n1 := &types.NodeInfoWrapper{Id: 1}
		n2 := &types.NodeInfoWrapper{Id: 2}
		n3 := &types.NodeInfoWrapper{Id: 3}
		n4 := &types.NodeInfoWrapper{Id: 4}
		n5 := &types.NodeInfoWrapper{Id: 1}

		s1 := []*types.NodeInfoWrapper{n1, n2, n3}
		s2 := []*types.NodeInfoWrapper{n2, n4, n5}

		got := PatchNeighbors(s1, s2)
		want := []*types.NodeInfoWrapper{n1, n2, n3, n4, n5}

		if !comp(want, got) {
			t.Errorf("1 common node: chooseMonitorTarget = %v, want %v", got, want)
		}
	}()

	func() {
		// all common nodes
		n1 := &types.NodeInfoWrapper{Id: 1}
		n2 := &types.NodeInfoWrapper{Id: 2}
		n3 := &types.NodeInfoWrapper{Id: 3}

		s1 := []*types.NodeInfoWrapper{n1, n2, n3}
		s2 := []*types.NodeInfoWrapper{n1, n2, n3}

		got := PatchNeighbors(s1, s2)
		want := []*types.NodeInfoWrapper{n1, n2, n3}

		if !comp(want, got) {
			t.Errorf("all common nodes: chooseMonitorTarget = %v, want %v", got, want)
		}
	}()
}
