package main

import (
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/utils"
	"log"
	"time"
)

const cloudHostName string = "cloud"
const testDuration = 1 * time.Minute
const testInterval = (1000 / 24) * time.Millisecond

func main() {
	timeout := time.After(testDuration)
	ticker := time.NewTicker(testInterval)
loop:
	for {
		select {
		case <-ticker.C:
			go func() {
				s := time.Now()
				nodes, err := core.Discover(1)
				if err != nil {
					log.Printf("failed to discover: %v", err)
					return
				}
				host := chooseHost(nodes, 0)

				consumer, err := core.Connect(host, core.DEFAULT_PORT)

				if err != nil {
					log.Fatalf("failed to connec to the server: %v", err)
				}

				call(&consumer, &types.CallRequest{
					AppId: 1,
					Body:  []byte{},
				})
				e := time.Since(s)
				log.Printf("overall: %s", e)
			}()
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
