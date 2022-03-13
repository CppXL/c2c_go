package server

import (
	"c2c/common/logger"
	"net"
)

type ServerConn struct {
	conn net.Conn
}

func (s *ServerConn) SetConn(conn net.Conn) {
	s.conn = conn
}

func (s *ServerConn) ReadData(n int) []byte {
	var data []byte
	for i := 0; i < n; i++ {
		_, err := s.conn.Read(data)
		logger.FatalIfError(err)

	}
}
