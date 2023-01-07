package main

import "log"

func main() {
	node := Node{}

	log.Fatal(Listen(node, DEFAULT_HOST, DEFAULT_PORT))
}
