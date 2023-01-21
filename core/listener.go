package core

import (
	"fmt"
	"net"
)

// DefectListener is wrapper of net.TCPLister.
//
// DefectListener sometimes fails to accept connection.
type DefectListener struct {
	*net.TCPListener
}

var count = 0

func (l *DefectListener) Accept() (net.Conn, error) {
	count += 1
	ll := l.TCPListener
	if count%3 == 0 {
		ll.Close()
		return nil, fmt.Errorf("connection closed by failure")
	}

	return ll.Accept()
}
