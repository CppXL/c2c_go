package connutil

import (
	"c2c/common/logger"
	"net"
)

func AcConn(ln net.Listener, handleConn func(conn net.Conn)) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		logger.Infof("Accept Conn from %v to %v \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleConn(conn)
	}
}
