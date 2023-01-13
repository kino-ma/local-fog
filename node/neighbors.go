package main

import (
	"local-fog/core/types"
)

var Neighbors []*types.NodeInfoWrapper

func UpdateNeighbors(neighbors []*types.NodeInfoWrapper) {
	Neighbors = neighbors
}
