package main

import (
	"fmt"
	"local-fog/core"
	"net"
	"os"

	"github.com/hashicorp/mdns"
)

const (
	serviceName = "_localfog._tcp"
	serviceTxt  = "LocalFog"
)

var serviceIp = net.IPv4(127, 0, 0, 1)

func RegisterAndServeMdns() error {
	host, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to get the hostname: %v", err)
	}

	info := []string{serviceTxt}
	ips := []net.IP{serviceIp}
	service, err := mdns.NewMDNSService(host, serviceName, "", "", core.DEFAULT_PORT, ips, info)

	if err != nil {
		return fmt.Errorf("failed ot create ne wmDNS service: %v", err)
	}

	_, err = mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return fmt.Errorf("failed to create new mdns server: %v", err)
	}

	return nil
}
