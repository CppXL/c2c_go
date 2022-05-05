package tcpagent

import (
	"c2c/common/logger"
	"net"
)

func (a *Agent) run() error {
	logger.Infof("tcp agent try to run")
	// try to read listen address and listen port from config

	return nil
}

func listen(a *Agent, network, address string) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		logger.Errorf(err.Error())
	}
	a.FrontListener = listener

	return nil
}
