package main

import (
	"local-fog/core"
	"local-fog/core/types"
	"log"
)

func main() {
	addr := Discover()
	log.Println(addr)
	consumer, err := core.Connect(addr.String(), core.DEFAULT_PORT)

	if err != nil {
		log.Fatalf("failed to connec to the server: %e", err)
	}

	consumer.Ping(&types.PingRequest{})
	consumer.Sync(&types.SyncRequest{})
	consumer.Call(&types.CallRequest{})
	consumer.GetProgram(&types.GetProgramRequest{})
}
