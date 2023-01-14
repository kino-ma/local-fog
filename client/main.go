package main

import (
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/utils"
	"log"
)

func main() {
	nodes, err := core.Discover(1)
	if err != nil {
		log.Fatalf("failed to discover: %v", err)
	}
	node := nodes[0]
	addr := utils.Uint32ToIp((node.AddrV4))

	log.Printf("discovered: %+v", addr)
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

	ur, err := consumer.UpdateNode(&types.UpdateNodeRequest{})
	if err != nil {
		log.Printf("UpdateNode request failed: %v", err)
	} else {
		log.Printf("UpdateNode: %v", ur)
	}
}
