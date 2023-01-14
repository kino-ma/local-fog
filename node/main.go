package main

import (
	"fmt"
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/utils"
	"log"
)

var info *types.NodeInfoWrapper

func main() {
	node := &Node{}

	neighbors, err := core.Discover(16)
	if err != nil {
		err = fmt.Errorf("failed to discover neighbors: %v", err)
		log.Fatal(err)
	}
	UpdateNeighbors(neighbors)

	addr, err := getPrimaryIp()
	if err != nil {
		err = fmt.Errorf("failed to get primary ip: %w", err)
		log.Fatal(err)
	}
	nodeId := utils.IpToUint32(addr)

	info = &types.NodeInfoWrapper{
		Id:     uint64(nodeId),
		AddrV4: utils.IpToUint32(addr),
	}
	log.Printf("Staring node %v", info)

	InsertNeighbor(info)
	log.Printf("neighbors including self: %v", Neighbors)

	organizer = chooseOrganizer(Neighbors)
	iAmOrganizer = organizer.Id == info.Id
	if iAmOrganizer {
		log.Print("I am the organizer")
	} else {
		log.Printf("node [%v] is the organizer", organizer.Id)
	}
	go ContinuosDiscovery()

	err = RegisterAndServeMdns(uint64(nodeId), addr)
	if err != nil {
		log.Fatalf("could not start mdns server: %v", err)
	}

	log.Fatal(core.Listen(node, core.DEFAULT_HOST, core.DEFAULT_PORT))
}
