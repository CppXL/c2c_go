package agent

import (
	"c2c/common/logger"
	"net"
	"sync"
)

// 这里
type Agent struct {
	conn          net.Conn
	AgentListener net.Listener
	ClientList    sync.Map
}

func (s *Agent) SetConn(c net.Conn) {
	s.conn = c
}
func (s *Agent) GetConn() net.Conn {
	return s.conn
}

func Listen(network, address string) (*Agent, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		logger.Errorf(err.Error())
	}

	return &Agent{AgentListener: listener}, nil
}

func (s *Agent) Run() {

}
