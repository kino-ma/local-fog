package main

import (
	"fmt"
	"local-fog/core"
	"local-fog/core/types"
	"log"
)

var NodeInfo *types.NodeInfo

func main() {
	node := &Node{}

	neighbors, err := core.Discover(16)
	if err != nil {
		err = fmt.Errorf("failed to discover neighbors: %v", err)
		log.Fatal(err)
	}

	nodeId := len(Neighbors) + 1
	addr, err := getPrimaryIp()
	if err != nil {
		err = fmt.Errorf("failed to get primary ip: %w", err)
		log.Fatal(err)
	}

	NodeInfo = &types.NodeInfo{
		Id:     uint64(nodeId),
		AddrV4: core.IpToUint32(addr),
	}

	UpdateNeighbors(neighbors)
	log.Print(Neighbors)

	err = RegisterAndServeMdns(uint64(nodeId), addr)
	if err != nil {
		log.Fatalf("could not start mdns server: %v", err)
	}

	log.Fatal(core.Listen(node, core.DEFAULT_HOST, core.DEFAULT_PORT))
}
