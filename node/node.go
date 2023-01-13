package main

import (
	"context"
	"fmt"

	"local-fog/core/apps"
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

	appId := t.AppId(p.AppId)
	body := p.Body

	out, err := apps.RunApp(appId, body)

	reply := t.CallReply{}
	reply.Output = out

	return &reply, err
}

func (n *Node) GetProgram(ctx context.Context, p *t.GetProgramRequest) (*t.GetProgramReply, error) {
	fmt.Printf("getProgram: id = %v\n", p.AppId)

	return &t.GetProgramReply{}, nil
}
