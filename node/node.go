package main

import (
	"context"
	"fmt"
	"log"

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
	log.Printf("before: %v", Neighbors)

	nodes := p.Nodes

	for _, n := range nodes {
		nn := (*t.NodeInfoWrapper)(n)
		InsertNeighbor(nn)
	}

	log.Printf("after: %v", Neighbors)

	outNodes := t.UnwrapNodeInfos(Neighbors)

	return &t.SyncReply{
		Nodes: outNodes,
	}, nil
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

func (n *Node) UpdateNode(ctx context.Context, p *t.UpdateNodeRequest) (*t.UpdateNodeReply, error) {
	log.Printf("updateNode: id = %+v, state = %v", p.Node, p.State)

	if p.Node == nil {
		err := fmt.Errorf("parameter 'Node' is nil")
		log.Printf("[ERROR] %v", err)
		return nil, err
	}

	switch p.State {
	case t.NodeState_JOINED:
		nn := (*t.NodeInfoWrapper)(p.Node)
		InsertNeighbor(nn)
	case t.NodeState_LEFT:
		nn := (*t.NodeInfoWrapper)(p.Node)
		err := DeleteNeighbor(nn)
		if err != nil {
			log.Printf("[WARN] Deletion of node [%v] was requested, but not found", nn)
		}
	}

	return &t.UpdateNodeReply{}, nil
}
