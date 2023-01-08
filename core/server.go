package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	t "local-fog/core/types"

	"google.golang.org/protobuf/proto"
)

const DEFAULT_HOST = "0.0.0.0"
const DEFAULT_PORT = 46866
const CONNECTION_TYPE = "tcp"

var logger = log.New(os.Stderr, "", 0)

type Server interface {
	HandlePing(p *t.Ping)
	HandleSync(s *t.Sync)
	HandleCall(c *t.Call)
	HandleGetProgram(g *t.GetProgram)
}

func Listen(s Server, host string, port int) (err error) {
	addr := host + ":" + fmt.Sprint(port)
	l, err := net.Listen(CONNECTION_TYPE, addr)

	if err != nil {
		err = fmt.Errorf("failed to linten on address %v: %e", addr, err)
		return err
	}

	logger.Printf("Start listening on '%v'...\n", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			err = fmt.Errorf("failed to accept connection: %e", err)
			return err
		}

		handle(&s, conn)
	}
}

func handle(s *Server, conn net.Conn) {
	defer conn.Close()
	ss := *s

	req := &t.Request{}
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		logger.Printf("Failed to read all from conn: %e\n", err)
		return
	}

	err = proto.Unmarshal(buf, req)
	if err != nil {
		logger.Printf("Failed to unmarshal request: %e\n", err)
		return
	}

	switch req.RequestType {
	case t.Request_PING:
		p := req.GetPing()

		if p == nil {
			typeName := t.Request_Type_name[int32(req.RequestType)]
			content := req.GetConetnt()
			logger.Printf("Requet type and content mismatch: type '%v (%v)' for %+v", typeName, req.RequestType, content)
			return
		}

		ss.HandlePing(p)

	case t.TYPE_SYNC:
		s := req.GetSync()

		if s == nil {
			typeName := t.Request_Type_name[int32(req.RequestType)]
			content := req.GetConetnt()
			logger.Printf("Requet type and content mismatch: type '%v (%v)' for %+v", typeName, req.RequestType, content)
			return
		}

		ss.HandleSync(s)

	case t.TYPE_CALL:
		c := req.GetCall()

		if c == nil {
			typeName := t.Request_Type_name[int32(req.RequestType)]
			content := req.GetConetnt()
			logger.Printf("Requet type and content mismatch: type '%v' for %+v", typeName, content)
			return
		}

		ss.HandleCall(c)

	case t.TYPE_GET_PROGRAM:
		g := req.GetGetProgram()

		if g == nil {
			typeName := t.Request_Type_name[int32(req.RequestType)]
			content := req.GetConetnt()
			logger.Printf("Requet type and content mismatch: type '%v' for %+v", typeName, content)
			return
		}

		ss.HandleGetProgram(g)
	}
}
