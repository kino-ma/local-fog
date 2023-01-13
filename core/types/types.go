package types

import (
	"fmt"
	"local-fog/core/utils"
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
	return fmt.Sprintf("node id:%v addr_v4:%v", i.Id, addr)
}
