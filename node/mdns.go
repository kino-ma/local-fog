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
	serviceTxt  = "v=localfog id=123"
)

var serviceIp net.IP
var internetHost = net.IP{8, 8, 8, 8}

func RegisterAndServeMdns() error {
	serviceIp, err := getPrimaryIp()
	if err != nil {
		return fmt.Errorf("failed to get primary ip: %w", err)
	}

	host, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to get the hostname: %v", err)
	}

	info := []string{serviceTxt}
	ips := []net.IP{serviceIp}
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
