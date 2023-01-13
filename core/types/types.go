package types

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
