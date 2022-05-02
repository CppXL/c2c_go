package client

import "net"

// 描述agent和client之间的连接，对于每个连接到agent的client来说都会有这样的结构存储相关信息
type Client struct {
	Token     uint32
	conn      net.Conn
	cryptInfo CryptBody
}

// 描述每个agent和client的连接加密部分信息+
type CryptBody struct {
	CryptType int
	SubType   int
	AESKey    []byte
	AESIv     []byte
}

func (c *Client) SetConn(conn net.Conn) {
	c.conn = conn
}
func (c *Client) GetConn() net.Conn {
	return c.conn
}

// 拨号函数
// func Dial(network, address string) (*Client, error) {

// }
