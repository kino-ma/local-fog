package main

import "net"

type RequestType int
type AppId uint64

const (
	TYPE_PING = iota // 0
	TYPE_SYNC
	TYPE_CALL
	TYPE_GET_PROGRAM
)

type Request struct {
	conn net.Conn
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
