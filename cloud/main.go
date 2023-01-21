package main

import (
	"local-fog/core"
	"log"
)

func main() {
	cloud := &Cloud{}
	log.Fatal(core.Listen(cloud, core.DEFAULT_HOST, core.DEFAULT_PORT))
}
