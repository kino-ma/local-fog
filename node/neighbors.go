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

func PeriodicTask() {
	ticker := time.NewTicker(syncPeriod)
	for range ticker.C {
		log.Print("ticker")
		organizerDiscovery()
		pingTarget()
	}
}

func organizerDiscovery() {
	if iAmOrganizer {
		log.Print("I am organizer. Running discovery...")
		nodes, err := core.Discover(16)
		if err != nil {
			err = fmt.Errorf("failed to discover: %w", err)
			log.Printf("[ERR] %v", err)
		}

		PatchNeighbors(nodes)

		organizer = chooseOrganizer(Neighbors)
		iAmOrganizer = organizer.Id == info.Id
	}
}

func pingTarget() {
	target := chooseMonitorTarget(Neighbors, info.Id)

	addr := utils.Uint32ToIp((target.AddrV4))
	consumer, err := core.Connect(addr.String(), core.DEFAULT_PORT)

	if err != nil {
		log.Printf("[ERROR] pingTarget: failed to connect to the server: %v", err)
		return
	}

	pr, err := consumer.Ping(&types.PingRequest{})
	if err != nil {
		log.Printf("[ERROR] Ping request failed: %v", err)
	} else {
		log.Printf("pingTarget success: %v", pr)
	}
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

func chooseMonitorTarget(ns []*types.NodeInfoWrapper, selfId uint64) *types.NodeInfoWrapper {
	self := &types.NodeInfoWrapper{Id: selfId}
	selfIdx, _ := types.FindNode(Neighbors, self)

	targetIdx := selfIdx - 1
	if targetIdx < 0 {
		targetIdx = len(Neighbors) - 1
	}

	return ns[targetIdx]
}
