package main

import (
	"c2c/common/logger"
	"c2c/common/utils/connutil"
	config "c2c/config/tcp"
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
	// 设置CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 解析命令行参数
	err := config.InitSrvConfig("/home/ub/code/c2c/c2c_go/server.yml")

	logger.FatalIfError(err)

	// 没错误就绑定端口

}
func main() {
	var err error
	// 这里要保证IP地址和端口是有效的
	// 监听配置中的Server端口
	ListenLn, err = net.Listen("tcp", config.SConfig.App.Agent.ListenAddr+":"+strconv.Itoa(int(config.SConfig.App.Agent.ListenPort)))
	logger.FatalIfError(err)
	if err != nil {
		fmt.Println("error lisning:", config.SConfig.App.Agent.ListenAddr+strconv.Itoa(int(config.SConfig.App.Agent.ListenPort)))
		log.Fatal(err)
	}
	// 监听control 端口
	CtrlLn, err = net.Listen("tcp", config.SConfig.App.Agent.ControlAddr+":"+strconv.Itoa(int(config.SConfig.App.Agent.ControlPort)))
	if err != nil {
		fmt.Println("error lisning:", config.SConfig.App.Agent.ControlAddr+strconv.Itoa(int(config.SConfig.App.Agent.ControlPort)))
		log.Fatal(err)
	}
	logger.Infof("Success listening %s:%d\t%s:%d\n",
		config.SConfig.App.Agent.ListenAddr, config.SConfig.App.Agent.ListenPort,
		config.SConfig.App.Agent.ControlAddr, config.SConfig.App.Agent.ControlPort)
	log.Printf("Begian Listen \n")

	// 起协程 使用Listen 返回的Listener类型的连接
	// 到这里就完成了监听任务
	go connutil.AcConn(ListenLn, command.HandleClientConn)
	connutil.AcConn(CtrlLn, command.HandleControlConn)

	// 阻塞
}
