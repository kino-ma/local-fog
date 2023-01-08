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
	HandlePing(p t.Ping)
	HandleSync(s t.Sync)
	HandleCall(c t.Call)
	HandleGetProgram(g t.GetProgram)
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

	typ, err := ReadByte(conn)

	if err != nil {
		logger.Printf("Failed to read 1 byte from connection: %v\n", err)
		return
	}

	req := t.Request{Conn: conn}

	switch typ {
	case t.TYPE_PING:
		p := t.Ping{
			Request: req,
		}
		ss.HandlePing(p)
	case t.TYPE_SYNC:
		sy := t.Sync{
			Request: req,
		}
		ss.HandleSync(sy)
	case t.TYPE_CALL:
		appId, err := Read8BytesNE(conn)

		if err != nil {
			logger.Printf("Failed to read AppId from connection: %v\n", err)
			return
		}

		body, err := ioutil.ReadAll(conn)

		if err != nil {
			logger.Printf("Failed to read Call body from connection: %v\n", err)
			return
		}

		c := t.Call{
			Request: req,
			AppId:   t.AppId(appId),
			Body:    body,
		}

		ss.HandleCall(c)

	case t.TYPE_GET_PROGRAM:
		appId, err := Read8BytesNE(conn)

		if err != nil {
			logger.Printf("Failed to read AppId from connection: %v\n", err)
			return
		}

		g := t.GetProgram{
			Request: req,
			AppId:   t.AppId(appId),
		}

		ss.HandleGetProgram(g)
	}
}
