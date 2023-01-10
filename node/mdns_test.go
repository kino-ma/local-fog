package main

import (
	"context"
	"testing"
)

func BenchmarkRegisterAndServeMdns10000(b *testing.B) {
	go RegisterAndServeMdns()

	conn := <-GetMdnsCann

	for i := 0; i < 10000; i++ {
		conn.Query(context.TODO(), "_localfog._tcp.local")
	}
}
