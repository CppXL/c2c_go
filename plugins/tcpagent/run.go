package tcpagent

import (
	"c2c/common/logger"
	config "c2c/configs"
	"fmt"
	"net"
	"strconv"
)

func (a *Agent) run() error {
	config.InitConfig("")

	logger.Infof("tcp agent try to run")
	// try to read listen address and listen port from config
	listenaddr := config.SConfig.Tcp.Agent.ListenAddr
	listenport := config.SConfig.Tcp.Agent.ListenPort
	err := listen(a, "tcp", listenaddr+":"+strconv.Itoa(int(listenport)))
	if err != nil {
		logger.Errorf(err.Error())
	} else {
		logger.Infof("tcp agent run success listening %s:%d\n", listenaddr, listenport)
		fmt.Println(listenaddr)
		fmt.Println(listenport)
	}
	return err
}

func listen(a *Agent, network, address string) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		logger.Errorf(err.Error())
	}
	a.FrontListener = listener

	return nil
}
