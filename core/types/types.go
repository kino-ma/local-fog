package types

import (
	"fmt"
	"local-fog/core/utils"

	"golang.org/x/exp/slices"
)

type RequestType int
type AppId uint64

const (
	_         = iota
	TYPE_PING = iota // 1
	TYPE_SYNC
	TYPE_CALL
	TYPE_GET_PROGRAM
)

type AppFunction func(body []byte) ([]byte, error)

type NodeInfoWrapper NodeInfo

func (i *NodeInfoWrapper) String() string {
	addr := utils.Uint32ToIp(i.AddrV4)
	return fmt.Sprintf("node id:%x addr_v4:%v", i.Id, addr)
}

func CompareNode(n1, n2 *NodeInfoWrapper) int {
	if n1.Id < n2.Id {
		return -1
	} else if n1.Id == n2.Id {
		return 0
	} else {
		return 1
	}
}

func FindNode(ns []*NodeInfoWrapper, n *NodeInfoWrapper) (int, bool) {
	return slices.BinarySearchFunc(ns, n, CompareNode)
}

func UnwrapNodeInfos(nodes []*NodeInfoWrapper) []*NodeInfo {
	out := make([]*NodeInfo, len(nodes))

	for i, n := range nodes {
		out[i] = (*NodeInfo)(n)
	}

	return out
}

func WrapNodeInfos(nodes []*NodeInfo) []*NodeInfoWrapper {
	out := make([]*NodeInfoWrapper, len(nodes))

	for i, n := range nodes {
		out[i] = (*NodeInfoWrapper)(n)
	}

	return out
}
