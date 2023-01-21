package main

import (
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/utils"
	"log"
	"time"
)

const cloudHostName string = "cloud"

func main() {
	nodes, err := core.Discover(1)
	if err != nil {
		log.Fatalf("failed to discover: %v", err)
	}
	var host string

	if len(nodes) < 1 {
		host = cloudHostName
	} else {
		node := nodes[0]
		addr := utils.Uint32ToIp((node.AddrV4))
		log.Printf("discovered: %+v", addr)
		host = addr.String()
	}

	consumer, err := core.Connect(host, core.DEFAULT_PORT)

	if err != nil {
		log.Fatalf("failed to connec to the server: %v", err)
	}

	s1 := time.Now()
	pr, err := consumer.Ping(&types.PingRequest{})
	if err != nil {
		log.Printf("Ping request failed: %v", err)
	} else {
		log.Printf("Ping: %v", pr)
		e1 := time.Since(s1)
		log.Printf("took %s", e1)
	}

	s2 := time.Now()
	sr, err := consumer.Sync(&types.SyncRequest{})
	if err != nil {
		log.Printf("Sync request failed: %v", err)
	} else {
		log.Printf("Sync: %v", sr)
		e2 := time.Since(s2)
		log.Printf("took %s", e2)
	}

	s3 := time.Now()
	cr, err := consumer.Call(&types.CallRequest{
		AppId: 1,
		Body:  []byte{},
	})
	if err != nil {
		log.Printf("Call request failed: %v", err)
	} else {
		log.Printf("Call: %v", cr)
		e3 := time.Since(s3)
		log.Printf("took %s", e3)
	}

	s4 := time.Now()
	gr, err := consumer.GetProgram(&types.GetProgramRequest{})
	if err != nil {
		log.Printf("GetProgram request failed: %v", err)
	} else {
		log.Printf("GetProgram: %v", gr)
		e4 := time.Since(s4)
		log.Printf("took %s", e4)
	}

	s5 := time.Now()
	ur, err := consumer.UpdateNode(&types.UpdateNodeRequest{})
	if err != nil {
		log.Printf("UpdateNode request failed: %v", err)
	} else {
		log.Printf("UpdateNode: %v", ur)
		e5 := time.Since(s5)
		log.Printf("took %s", e5)
	}
}
