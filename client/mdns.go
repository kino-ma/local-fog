package main

import (
	"log"
	"net"

	"github.com/hashicorp/mdns"
)

func Discover() (addr net.IP) {
	ch := make(chan *mdns.ServiceEntry, 4)

	// var entry *mdns.ServiceEntry
	i := 0

	go func() {
		for entry := range ch {
			log.Printf("got entry[%v]: %+v", i, entry)
			i += 1
		}
	}()

	mdns.Lookup("_localfog._tcp", ch)
	close(ch)

	return
}
