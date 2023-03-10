package main

import (
	"fmt"
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/types/helper"
	"local-fog/core/utils"
	"log"
	"sort"
	"time"
)

const syncPeriod = 5 * time.Second

// Neighbors contains neighbors. Note: Many functions assume this slice to be sorted. Do not modify directly.
var Neighbors []*types.NodeInfoWrapper

var (
	ErrNeighborNotFound = fmt.Errorf("neighbor with that id was not found")
)

// UpdateNeighbors overwrites neighbors list by given one.
// Note: this function sorts the argument slice first, i.e., breaks original order.
func UpdateNeighbors(neighbors []*types.NodeInfoWrapper) bool {
	sortNeighbors(neighbors)
	changed := !types.IsSameList(neighbors, Neighbors)

	Neighbors = neighbors

	return changed
}

func PatchNeighbors(neighbors []*types.NodeInfoWrapper) bool {
	sortNeighbors(neighbors)
	ns := patchNodes(Neighbors, neighbors)
	return UpdateNeighbors(ns)
}

func InsertNeighbor(neigh *types.NodeInfoWrapper) {
	Neighbors, _ = utils.InsertSorted(Neighbors, neigh, types.CompareNode)
}

func DeleteNeighbor(neigh *types.NodeInfoWrapper) error {
	i, found := types.FindNode(Neighbors, neigh)
	if !found {
		return ErrNeighborNotFound
	}

	Neighbors = utils.RemoveIndex(Neighbors, i)
	return nil
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
	organizer := chooseOrganizer(Neighbors)
	iAmOrganizer := organizer.Id == info.Id

	log.Printf("Self = [%v] , my list = %v , organizer = [%v]", info, Neighbors, organizer)

	if iAmOrganizer {
		log.Print("I am organizer. Running discovery...")
		nodes, err := core.Discover(16)
		if err != nil {
			err = fmt.Errorf("failed to discover: %w", err)
			log.Printf("[ERR] %v", err)
		}

		hasChanged := PatchNeighbors(nodes)

		organizer = chooseOrganizer(Neighbors)
		iAmOrganizer = organizer.Id == info.Id

		if hasChanged && iAmOrganizer {
			ns := types.RemoveNode(Neighbors, info)
			err := syncAll(ns)
			if err != nil {
				log.Printf("[ERROR] %v", err)
			}
		}
	}
}

func pingTarget() {
	target := chooseMonitorTarget(Neighbors, info.Id)

	if target == nil {
		log.Printf("[INFO] no need to monitor. Skipping")
		return
	}

	addr := utils.Uint32ToIp((target.AddrV4))
	consumer, err := core.Connect(addr.String(), core.DEFAULT_PORT)

	if err != nil {
		log.Printf("[ERROR] pingTarget: failed to connect to the server: %v", err)
		return
	}

	_, err = consumer.Ping(&types.PingRequest{})
	if err == nil {
		log.Printf("pingTarget success: %v", target)
		return
	}

	// if ping failed, remove its information from all node

	log.Printf("[ERROR] Ping request failed: %v", err)
	log.Printf("start removing node [%x]", target.Id)

	DeleteNeighbor(target)
	ns := types.RemoveNode(Neighbors, info)
	err = deleteFromAll(ns, target)
	if err != nil {
		log.Printf("[ERROR] failed to delete node from all nodes: %v", err)
		return
	}
}

func syncAll(ns []*types.NodeInfoWrapper) error {
	syncReq := func(n *types.NodeInfoWrapper, consumer core.FogConsumer) error {
		nodesToSend := types.UnwrapNodeInfos(Neighbors)

		sReq := &types.SyncRequest{
			Nodes: nodesToSend,
		}

		sRep, err := consumer.Sync(sReq)
		if err != nil {
			err = fmt.Errorf("anonymous sync: failed to sync with node [%v]: %w", n.Id, err)
			return err
		}

		nodes := types.WrapNodeInfos(sRep.Nodes)
		PatchNeighbors(nodes)

		log.Printf("synced with node [%v]", n.Id)
		return nil
	}

	err := helper.RequestForAllNode(ns, syncReq)

	if err != nil {
		return fmt.Errorf("syncAll: 1 ore more errors occured while syncing: %v", err)
	}

	return nil
}

func addToAll(ns []*types.NodeInfoWrapper, target *types.NodeInfoWrapper) error {
	updateReq := func(dest *types.NodeInfoWrapper, consumer core.FogConsumer) error {
		node := (*types.NodeInfo)(target)

		uReq := &types.UpdateNodeRequest{
			Node:  node,
			State: types.NodeState_JOINED,
		}

		_, err := consumer.UpdateNode(uReq)
		if err != nil {
			err = fmt.Errorf("anonymous updateNode: failed to update information of node [%x]: %w", target.Id, err)
			return err
		}

		log.Printf("added [%v] to node [%v]", target, dest)

		return nil
	}

	err := helper.RequestForAllNode(ns, updateReq)

	if err != nil {
		return fmt.Errorf("deleteFromAll: 1 ore more errors occured while syncing: %w", err)
	}

	return nil
}

func deleteFromAll(ns []*types.NodeInfoWrapper, target *types.NodeInfoWrapper) error {
	updateReq := func(dest *types.NodeInfoWrapper, consumer core.FogConsumer) error {
		node := (*types.NodeInfo)(target)

		uReq := &types.UpdateNodeRequest{
			Node:  node,
			State: types.NodeState_LEFT,
		}

		_, err := consumer.UpdateNode(uReq)
		if err != nil {
			err = fmt.Errorf("anonymous updateNode: failed to update information of node [%v]: %w", target, err)
			return err
		}
		return nil
	}

	err := helper.RequestForAllNode(ns, updateReq)

	if err != nil {
		return fmt.Errorf("deleteFromAll: 1 ore more errors occured while syncing: %w", err)
	}

	return nil
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
	if len(ns) == 1 {
		return nil
	}

	self := &types.NodeInfoWrapper{Id: selfId}
	selfIdx, _ := types.FindNode(Neighbors, self)

	targetIdx := selfIdx - 1
	if targetIdx < 0 {
		targetIdx = len(Neighbors) - 1
	}

	return ns[targetIdx]
}

func patchNodes(target, patch []*types.NodeInfoWrapper) []*types.NodeInfoWrapper {
	out := make([]*types.NodeInfoWrapper, len(target))
	copy(out, target)
	out = append(out, patch...)
	sortNeighbors(out)
	out = utils.RemoveDuplicates(out, types.CompareNode)
	return out
}
