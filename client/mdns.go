package main

import (
	"net"

	"github.com/hashicorp/mdns"
)

func Discover() net.IP {
	ch := make(chan *mdns.ServiceEntry, 1)
	mdns.Lookup("_localfog._tcp", ch)
	entry := <-ch
	close(ch)

	return entry.AddrV4
}
