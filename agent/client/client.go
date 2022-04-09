package client

import "net"

// 描述连接到agent的client
type Client struct {
	Token uint32
	conn  net.Conn
}

func (s *Client) SetConn(c net.Conn) {
	s.conn = c
}
func (s *Client) GetConn() net.Conn {
	return s.conn
}

func Dial(network, address string) (*Client, error) {

}
