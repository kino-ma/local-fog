package main

import (
	"context"
	"fmt"

	t "local-fog/core/types"
)

type Node struct {
	t.UnimplementedLocalFogServer
}

func (n *Node) Ping(ctx context.Context, p *t.PingRequest) (*t.PingReply, error) {
	fmt.Printf("ping\n")

	return &t.PingReply{}, nil
}

func (n *Node) Sync(ctx context.Context, p *t.SyncRequest) (*t.SyncReply, error) {
	fmt.Printf("sync\n")

	return &t.SyncReply{}, nil
}

func (n *Node) Call(ctx context.Context, p *t.CallRequest) (*t.CallReply, error) {
	fmt.Printf("call: id = %v, body = %v\n", p.AppId, p.Body)

	return &t.CallReply{}, nil
}

func (n *Node) GetProgram(ctx context.Context, p *t.GetProgramRequest) (*t.GetProgramReply, error) {
	fmt.Printf("getProgram: id = %v\n", p.AppId)

	return &t.GetProgramReply{}, nil
}
