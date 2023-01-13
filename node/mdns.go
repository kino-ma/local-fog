package main

import (
	"fmt"
	"local-fog/core"
	"log"
	"net"
	"os"

	"github.com/google/gopacket/routing"
	"github.com/hashicorp/mdns"
)

const (
	serviceName = "_localfog._tcp"
)

var internetHost = net.IP{8, 8, 8, 8}

func RegisterAndServeMdns(nodeId uint64, addr net.IP) error {
	host, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to get the hostname: %v", err)
	}

	serviceTxt := core.NewTxt(uint64(nodeId))
	info := []string{serviceTxt}
	ips := []net.IP{addr}

	service, err := mdns.NewMDNSService(host, serviceName, "local.", "", core.DEFAULT_PORT, ips, info)

	if err != nil {
		return fmt.Errorf("failed ot create ne wmDNS service: %v", err)
	}

	_, err = mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return fmt.Errorf("failed to create new mdns server: %v", err)
	}

	return nil
}

func getPrimaryIp() (net.IP, error) {
	router, err := routing.New()
	if err != nil {
		return nil, err
	}

	_, _, primaryIp, err := router.Route(internetHost)
	log.Printf("ip = %v, err = %v", primaryIp, err)
	return primaryIp, err
}
