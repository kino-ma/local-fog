package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/mdns"
)

func Discover() (net.IP, error) {
	ch := make(chan *mdns.ServiceEntry)

	err := mdns.Lookup("_localfog._tcp", ch)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup the service: %v", err)
	}

	log.Printf("start lookup")

	entry := <-ch
	log.Printf("got entry: %v", entry)
	close(ch)

	return entry.AddrV4, err
}
