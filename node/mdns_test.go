package main

import (
	"context"
	"testing"

	"github.com/hashicorp/mdns"
)

var conn *mdns.Conn

func init() {
	go RegisterAndServeMdns()
	conn = <-GetMdnsCann
}

func BenchmarkRegisterAndServeMdns10000(b *testing.B) {
	_, _, err := conn.Query(context.TODO(), "_localfog._tcp.local")

	if err != nil {
		b.Error(err)
	}
}
