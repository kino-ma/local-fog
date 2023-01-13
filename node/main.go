package main

import (
	"fmt"
	"local-fog/core"
	"log"
)

func main() {
	node := &Node{}

	neighbors, err := core.Discover(16)
	if err != nil {
		err = fmt.Errorf("failed to discover neighbors: %v", err)
		log.Fatal(err)
	}

	UpdateNeighbors(neighbors)
	log.Print(Neighbors)

	err = RegisterAndServeMdns()
	if err != nil {
		log.Fatalf("could not start mdns server: %v", err)
	}

	log.Fatal(core.Listen(node, core.DEFAULT_HOST, core.DEFAULT_PORT))
}
