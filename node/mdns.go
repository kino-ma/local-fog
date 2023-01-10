package main

import (
	"fmt"
	"net"

	"github.com/pion/mdns"
	"golang.org/x/net/ipv4"
)

var GetMdnsCann chan *mdns.Conn = make(chan *mdns.Conn)

func RegisterAndServeMdns() error {
	addr, err := net.ResolveUDPAddr("udp", mdns.DefaultAddress)
	if err != nil {
		return fmt.Errorf("failed to resolve udp address: %v", err)
	}

	l, err := net.ListenUDP("udp4", addr)
	if err != nil {
		return fmt.Errorf("failed to listen udp: %v", err)
	}

	conn, err := mdns.Server(ipv4.NewPacketConn(l), &mdns.Config{
		LocalNames: []string{"_localfog._tcp.local"},
	})
	if err != nil {
		close(GetMdnsCann)
		return fmt.Errorf("failed to create mdns server: %v", err)
	}
	GetMdnsCann <- conn

	select {}
}
