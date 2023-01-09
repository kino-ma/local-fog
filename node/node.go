package main

import (
	"fmt"

	t "local-fog/core/types"
)

type Node struct {
	t.UnimplementedLocalFogServer
}

func (n *Node) Ping(p *t.PingRequest) (*t.PingReply, error) {
	fmt.Printf("ping: %+v\n", p)

	return &t.PingReply{}, nil
}

func (n *Node) Sync(p *t.SyncRequest) (*t.SyncReply, error) {
	fmt.Printf("ping: %+v\n", p)

	return &t.SyncReply{}, nil
}

func (n *Node) Call(p *t.CallRequest) (*t.CallReply, error) {
	fmt.Printf("ping: %+v\n", p)

	return &t.CallReply{}, nil
}

func (n *Node) GetProgram(p *t.GetProgramRequest) (*t.GetProgramReply, error) {
	fmt.Printf("ping: %+v\n", p)

	return &t.GetProgramReply{}, nil
}
