package main

import (
	"local-fog/core"
	"log"
)

func main() {
	node := &Node{}

	go func() {
		err := RegisterAndServeMdns()

		if err != nil {
			panic(err)
		}
	}()

	log.Fatal(core.Listen(node, core.DEFAULT_HOST, core.DEFAULT_PORT))
}
