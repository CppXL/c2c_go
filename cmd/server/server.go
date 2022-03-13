package main

import (
	"c2c/common/logger"
	"c2c/common/utils/connutil"
	"c2c/config"
	"c2c/control/command"
	"fmt"
	"log"
	"net"
	"runtime"
	"strconv"
)

var (
	// 监听端口和控制端口连接
	ListenLn, CtrlLn net.Listener
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 解析命令行参数
	var err error
	err = config.InitSrvConfig("/home/ub/code/c2c/c2c_go/server.yml")
	if err != nil {
		logger.Errorf(err.Error())
	}

	// 没错误就绑定端口

}
func main() {
	var err error
	ListenLn, err = net.Listen("tcp", config.SConfig.App.Server.ListenAddr+":"+strconv.Itoa(int(config.SConfig.App.Server.ListenPort)))
	if err != nil {
		fmt.Println(config.SConfig.App.Server.ListenAddr + strconv.Itoa(int(config.SConfig.App.Server.ListenPort)))
		log.Fatal(err)
	}
	CtrlLn, err = net.Listen("tcp", config.SConfig.App.Server.ControlAddr+":"+strconv.Itoa(int(config.SConfig.App.Server.ControlPort)))
	if err != nil {
		log.Fatal(err)
	}
	logger.Infof("Success listening %s:%d\t%s:%d\n",
		config.SConfig.App.Server.ListenAddr, config.SConfig.App.Server.ListenPort,
		config.SConfig.App.Server.ControlAddr, config.SConfig.App.Server.ControlPort)
	log.Printf("Begian Listen \n")

	go connutil.AcConn(ListenLn, command.HandleClientConn)
	connutil.AcConn(CtrlLn, command.HandleControlConn)

	// 阻塞
}
