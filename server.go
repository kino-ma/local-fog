package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const DEFAULT_HOST = "0.0.0.0"
const DEFAULT_PORT = "46866"
const CONNECTION_TYPE = "tcp"

var logger = log.New(os.Stderr, "", 0)

type Server interface {
	Handle(conn net.Conn)
}

func Listen(s Server, host string, port int) (err error) {
	addr := DEFAULT_HOST + ":" + DEFAULT_PORT
	l, err := net.Listen(CONNECTION_TYPE, addr)

	if err != nil {
		err = fmt.Errorf("failed to linten on address %v: %e", addr, err)
		return err
	}

	defer l.Close()
	logger.Printf("Start listening on '%v'...\n", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			err = fmt.Errorf("failed to accept connection: %e", err)
			return err
		}

		go s.Handle(conn)
	}
}
