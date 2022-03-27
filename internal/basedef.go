package internal

import "net"

type Baseconn struct {
	fd   net.Conn
	quit chan struct{}
}

func (s *Baseconn) SetConn(conn net.Conn) {
	s.fd = conn
}
func (s *Baseconn) ReadData(recv []byte, n int) {
}
func (s *Baseconn) Close() error {
	return s.fd.Close()

}
