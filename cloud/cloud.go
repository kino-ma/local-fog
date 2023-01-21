package main

import (
	"context"
	"fmt"
	"log"

	"local-fog/core/apps"
	t "local-fog/core/types"
)

type Cloud struct {
	t.UnimplementedLocalFogServer
}

var (
	ErrNotImplemented = fmt.Errorf("this request is not implemented for Cloud")
)

func (n *Cloud) Ping(ctx context.Context, p *t.PingRequest) (*t.PingReply, error) {
	fmt.Printf("ping\n")
	err := ErrNotImplemented
	log.Print(err)

	return nil, err
}

func (n *Cloud) Sync(ctx context.Context, p *t.SyncRequest) (*t.SyncReply, error) {
	fmt.Printf("sync\n")
	err := ErrNotImplemented
	log.Print(err)

	return nil, err
}

func (n *Cloud) Call(ctx context.Context, p *t.CallRequest) (*t.CallReply, error) {
	fmt.Printf("call: id = %v, body = %v\n", p.AppId, p.Body)

	appId := t.AppId(p.AppId)
	body := p.Body

	out, err := apps.RunApp(appId, body)

	reply := t.CallReply{}
	reply.Output = out

	return &reply, err
}

func (n *Cloud) GetProgram(ctx context.Context, p *t.GetProgramRequest) (*t.GetProgramReply, error) {
	fmt.Printf("getProgram: id = %v\n", p.AppId)

	err := ErrNotImplemented
	log.Print(err)

	return nil, err
}

func (n *Cloud) UpdateNode(ctx context.Context, p *t.UpdateNodeRequest) (*t.UpdateNodeReply, error) {
	log.Printf("updateNode: id = %+v, state = %v", p.Node, p.State)

	err := ErrNotImplemented
	log.Print(err)

	return nil, err
}
