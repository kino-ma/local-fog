package main

import (
	"fmt"
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/utils"
	"log"
	"sort"
	"time"
)

const syncPeriod = time.Minute

// Neighbors contains neighbors. Note: Many functions assume this slice to be sorted. Do not modify directly.
var Neighbors []*types.NodeInfoWrapper
var (
	organizer    *types.NodeInfoWrapper
	iAmOrganizer bool
)

var (
	ErrNeighborNotFound = fmt.Errorf("neighbor with that id was not found")
)

// UpdateNeighbors overwrites neighbors list by given one.
// Note: this function sorts the argument slice first, i.e., breaks original order.
func UpdateNeighbors(neighbors []*types.NodeInfoWrapper) {
	sortNeighbors(neighbors)
	Neighbors = neighbors
}

func PatchNeighbors(patch []*types.NodeInfoWrapper) {
	ns := append(Neighbors, patch...)
	sortNeighbors(ns)
	ns = utils.RemoveDuplicates(ns, types.CompareNode)
	Neighbors = ns
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

func ContinuosDiscovery() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		periodicSync()
	}
}

func periodicSync() {
	if iAmOrganizer {
		log.Print("I am organizer. Running discovery...")
		nodes, err := core.Discover(16)
		if err != nil {
			err = fmt.Errorf("failed to discover: %w", err)
			log.Printf("[ERR] %v", err)
		}

		PatchNeighbors(nodes)

		o := chooseOrganizer(Neighbors)
		iAmOrganizer = o.Id == info.Id
	}

	// do monitoring
}

func nodesXor(n1, n2 []*types.NodeInfoWrapper) []*types.NodeInfoWrapper {
	sortNeighbors(n1)
	sortNeighbors(n2)

	return utils.XorSlice(n1, n2, types.CompareNode)
}

func sortNeighbors(ns []*types.NodeInfoWrapper) {
	compareId := func(i, j int) bool { return ns[i].Id < ns[j].Id }
	sort.Slice(ns, compareId)
}

func chooseOrganizer(ns []*types.NodeInfoWrapper) *types.NodeInfoWrapper {
	return ns[0]
}
