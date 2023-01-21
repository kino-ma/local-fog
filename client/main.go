package main

import (
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/utils"
	"log"
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

	call(&consumer, &types.CallRequest{
		AppId: 1,
		Body:  []byte{},
	})
}
