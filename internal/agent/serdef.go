package agent

import (
	"c2c/internal"
	"net"
)

type ServerConn struct {
	baseconn internal.Baseconn
}

func (s *ServerConn) SetConn(conn net.Conn) {
	s.baseconn.SetConn(conn)
}

func (s *ServerConn) ReadN(n int) []byte {
	var data []byte
	for i := 0; i < n; i++ {
		s.baseconn.ReadData(data, 1)

	}
}
