package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const DEFAULT_HOST = "0.0.0.0"
const DEFAULT_PORT = 46866
const CONNECTION_TYPE = "tcp"

var logger = log.New(os.Stderr, "", 0)

type Server interface {
	HandlePing(p Ping)
	HandleSync(s Sync)
	HandleCall(c Call)
	HandleGetProgram(g GetProgram)
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

	typ, err := ReadByte(conn)

	if err != nil {
		logger.Printf("Failed to read 1 byte from connection: %v\n", err)
		return
	}

	req := Request{conn}

	switch typ {
	case TYPE_PING:
		p := Ping{
			Request: req,
		}
		ss.HandlePing(p)
	case TYPE_SYNC:
		sy := Sync{
			Request: req,
		}
		ss.HandleSync(sy)
	case TYPE_CALL:
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

		c := Call{
			Request: req,
			AppId:   AppId(appId),
			Body:    body,
		}

		ss.HandleCall(c)

	case TYPE_GET_PROGRAM:
		appId, err := Read8BytesNE(conn)

		if err != nil {
			logger.Printf("Failed to read AppId from connection: %v\n", err)
			return
		}

		g := GetProgram{
			Request: req,
			AppId:   AppId(appId),
		}

		ss.HandleGetProgram(g)
	}
}
