package main

import "net"

type NodeInfo struct {
	Addr net.IPAddr
	Id   uint64
}

var Neighbors []NodeInfo = []NodeInfo{}
