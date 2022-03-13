package client

import "net"

type ClientConn struct {
	conn net.Conn
	quit chan struct{}
}
