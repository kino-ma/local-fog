package core

import "net"

// DefectListener is wrapper of net.TCPLister.
//
// DefectListener sometimes fails to accept connection.
type DefectListener struct {
	net.TCPListener
}

func (l *DefectListener) Accept() (net.Conn, error) {
	ll := l.TCPListener
	ll.Close()
	return nil, net.ErrClosed
}
