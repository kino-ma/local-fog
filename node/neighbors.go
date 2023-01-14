package main

import (
	"fmt"
	"local-fog/core/types"
	"local-fog/core/utils"
	"sort"
)

// Neighbors contains neighbors. Note: Many functions assume this slice to be sorted. Do not modify directly.
var Neighbors []*types.NodeInfoWrapper

var (
	ErrNeighborNotFound = fmt.Errorf("neighbor with that id was not found")
)

// UpdateNeighbors overwrites neighbors list by given one.
// Note: this function sorts the argument slice first, i.e., breaks original order.
func UpdateNeighbors(neighbors []*types.NodeInfoWrapper) {
	sortNeighbors(neighbors)
	Neighbors = neighbors
}

func InsertNeighbor(neigh *types.NodeInfoWrapper) {
	Neighbors, _ = utils.InsertSorted(Neighbors, neigh, types.CompareNode)
}

func DeleteNeighbor(neigh *types.NodeInfoWrapper) error {
	for i, n := range Neighbors {
		if n.Id == neigh.Id {
			Neighbors = append(Neighbors[:i], Neighbors[i+1:]...)
			return nil
		}
	}

	return ErrNeighborNotFound
}

func sortNeighbors(ns []*types.NodeInfoWrapper) {
	compareId := func(i, j int) bool { return ns[i].Id < ns[i].Id }
	sort.Slice(ns, compareId)
}

func chooseOrganizer(ns []*types.NodeInfoWrapper) *types.NodeInfoWrapper {
	return ns[0]
}
