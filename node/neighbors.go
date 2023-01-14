package main

import (
	"fmt"
	"local-fog/core/types"
)

var Neighbors []*types.NodeInfoWrapper

var (
	ErrNeighborNotFound = fmt.Errorf("neighbor with that id was not found")
)

func UpdateNeighbors(neighbors []*types.NodeInfoWrapper) {
	Neighbors = neighbors
}

func InsertNeighbor(neigh *types.NodeInfoWrapper) {
	Neighbors = append(Neighbors, neigh)
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
