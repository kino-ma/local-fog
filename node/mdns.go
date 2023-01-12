package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/mdns"
)

func RegisterAndServeMdns() error {
	host, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to get the hostname: %v", err)
	}

	info := []string{"LocalFog"}
	ips := []net.IP{net.IPv4(127, 0, 0, 1)}
	service, err := mdns.NewMDNSService(host, "_localfog._tcp", "", "", 46866, ips, info)

	if err != nil {
		return fmt.Errorf("failed ot create ne wmDNS service: %v", err)
	}

	_, err = mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return fmt.Errorf("failed to create new mdns server: %v", err)
	}

	return nil
}
