package tcpagent

import (
	"c2c/common/logger"
	"net"
	"sync"
)

// 描述 agent自身的连接
type Agent struct {
	FrontListener net.Listener
	BackendConn   net.Conn
	ClientList    sync.Map
}

func Listen(network, address string) (*Agent, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		logger.Errorf(err.Error())
	}

	return &Agent{FrontListener: listener}, nil
}

func (s *Agent) Run() {

}
