package types

import "net"

type RequestType int
type AppId uint64

const (
	_         = iota
	TYPE_PING = iota // 1
	TYPE_SYNC
	TYPE_CALL
	TYPE_GET_PROGRAM
)

type Request struct {
	Conn net.Conn
}

type Ping struct {
	Request
}
type Sync struct {
	Request
}
type Call struct {
	Request
	AppId AppId
	Body  []byte
}
type GetProgram struct {
	Request
	AppId AppId
}
