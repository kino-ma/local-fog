package main

import (
	"local-fog/core"
	"log"
)

func main() {
	node := &Node{}

	err := RegisterAndServeMdns()
	if err != nil {
		log.Fatalf("could not start mdns server: %v", err)
	}

	log.Fatal(core.Listen(node, core.DEFAULT_HOST, core.DEFAULT_PORT))
}
