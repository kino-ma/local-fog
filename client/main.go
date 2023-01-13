package main

import (
	"local-fog/core"
	"local-fog/core/types"
	"log"
)

func main() {
	addr, err := Discover()
	if err != nil {
		log.Fatalf("failed to discover: %v", err)
	}

	log.Printf("discovered: %v", addr)
	consumer, err := core.Connect(addr.String(), core.DEFAULT_PORT)

	if err != nil {
		log.Fatalf("failed to connec to the server: %v", err)
	}

	pr, err := consumer.Ping(&types.PingRequest{})
	if err != nil {
		log.Printf("Ping request failed: %v", err)
	} else {
		log.Printf("Ping: %v", pr)
	}

	sr, err := consumer.Sync(&types.SyncRequest{})
	if err != nil {
		log.Printf("Sync request failed: %v", err)
	} else {
		log.Printf("Sync: %v", sr)
	}

	cr, err := consumer.Call(&types.CallRequest{
		AppId: 1,
		Body:  []byte{},
	})
	if err != nil {
		log.Printf("Call request failed: %v", err)
	} else {
		log.Printf("Call: %v", cr)
	}

	gr, err := consumer.GetProgram(&types.GetProgramRequest{})
	if err != nil {
		log.Printf("GetProgram request failed: %v", err)
	} else {
		log.Printf("GetProgram: %v", gr)
	}
}
