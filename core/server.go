package core

import (
	"fmt"
	"log"
	"net"
	"os"

	t "local-fog/core/types"

	"google.golang.org/grpc"
)

const DEFAULT_HOST = "0.0.0.0"
const DEFAULT_PORT = 46866
const CONNECTION_TYPE = "tcp"

var logger = log.New(os.Stderr, "", 0)

func Listen(s t.LocalFogServer, host string, port int) (err error) {
	addr := host + ":" + fmt.Sprint(port)
	l, err := net.Listen(CONNECTION_TYPE, addr)
	// tl, _ := l.(*net.TCPListener)
	// ll := &DefectListener{tl}

	if err != nil {
		err = fmt.Errorf("failed to linten on address %v: %e", addr, err)
		return err
	}

	gs := grpc.NewServer()
	t.RegisterLocalFogServer(gs, s)

	logger.Printf("Start listening on '%v'...\n", addr)

	// return gs.Serve(ll)
	return gs.Serve(l)
}
