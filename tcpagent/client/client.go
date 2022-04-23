package client

import "net"

// 描述连接到agent的client
type Client struct {
	Token     uint32
	conn      net.Conn
	cryptInfo CryptBody
}

type CryptBody struct {
	CryptType int
	Key       []byte
	Iv        []byte
}

func (s *Client) SetConn(c net.Conn) {
	s.conn = c
}
func (s *Client) GetConn() net.Conn {
	return s.conn
}

// 拨号函数
// func Dial(network, address string) (*Client, error) {

// }
