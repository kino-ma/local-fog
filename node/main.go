package main

import (
	"local-fog/core"
	"log"
)

func main() {
	node := &Node{}

	log.Fatal(core.Listen(node, core.DEFAULT_HOST, core.DEFAULT_PORT))
}
