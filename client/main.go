package main

import (
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/utils"
	"log"
	"time"
)

const cloudHostName string = "cloud"
const testDuration = 5 * time.Second
const testInterval = 1 * time.Millisecond

func main() {
	nodes, err := core.Discover(16)
	if err != nil {
		log.Fatalf("failed to discover: %v", err)
	}
	timeout := time.After(testDuration)
	ticker := time.NewTicker(testInterval)
	l := len(nodes)
	if l == 0 {
		l = 1
	}

loop:
	for i := 0; ; i = (i + 1) % l {
		select {
		case <-ticker.C:
			go func(i int) {
				host := chooseHost(nodes, i)
				consumer, err := core.Connect(host, core.DEFAULT_PORT)

				if err != nil {
					log.Fatalf("failed to connec to the server: %v", err)
				}

				call(&consumer, &types.CallRequest{
					AppId: 1,
					Body:  []byte{},
				})
			}(i)
		case <-timeout:
			break loop
		}
	}
}

func chooseHost(ns []*types.NodeInfoWrapper, i int) string {
	if len(ns) < i+1 {
		return cloudHostName
	}

	node := ns[i]
	addr := utils.Uint32ToIp((node.AddrV4))
	log.Printf("discovered: %+v", addr)
	return addr.String()
}
