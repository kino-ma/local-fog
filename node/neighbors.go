package main

import (
	"local-fog/core/types"
)

var Neighbors []types.NodeInfo = []types.NodeInfo{}

func UpdateNeighbors(neighbors []types.NodeInfo) {
	Neighbors = neighbors
}
