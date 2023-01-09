package main

import (
	"local-fog/core"
	"local-fog/core/types"
	"log"
)

func main() {
	consumer, err := core.Connect(core.DEFAULT_HOST, core.DEFAULT_PORT)

	if err != nil {
		log.Fatalf("failed to connec to the server: %e", err)
	}

	consumer.Ping(&types.PingRequest{})
	consumer.Sync(&types.SyncRequest{})
	consumer.Call(&types.CallRequest{})
	consumer.GetProgram(&types.GetProgramRequest{})
}
