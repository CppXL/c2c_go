package main

// import (
// 	"c2c/common/logger"
// 	"c2c/common/utils/connutil"
// 	"c2c/config"
// 	"c2c/tcpagent/heartbeat"
// 	"fmt"
// 	"log"
// 	"net"
// 	"runtime"
// 	"strconv"
// )

// var (
// 	// 监听端口和控制端口连接
// 	ListenLn, CtrlLn net.Listener
// )

// func init() {
// 	// set CPU
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// 	// paese command line args
// 	err := config.InitSrvConfig("/home/ub/code/c2c/c2c/server.yml")

// 	logger.FatalIfError(err)

// 	// 没错误就绑定端口

// }
// func main() {
// 	var err error
// 	// 这里要保证IP地址和端口是有效的
// 	// 监听配置中的Server端口
// 	ListenLn, err = net.Listen("tcp", config.SConfig.Tcp.Agent.ListenAddr+":"+
// 		strconv.Itoa(int(config.SConfig.Tcp.Agent.ListenPort)))
// 	logger.FatalIfError(err)
// 	if err != nil {
// 		fmt.Println("error lisning:", config.SConfig.Tcp.Agent.ListenAddr+
// 			strconv.Itoa(int(config.SConfig.Tcp.Agent.ListenPort)))
// 		log.Fatal(err)
// 	}
// 	// 监听control 端口
// 	CtrlLn, err = net.Listen("tcp", config.SConfig.Tcp.Agent.ControlAddr+":"+
// 		strconv.Itoa(int(config.SConfig.Tcp.Agent.ControlPort)))
// 	if err != nil {
// 		fmt.Println("error lisning:", config.SConfig.Tcp.Agent.ControlAddr+
// 			strconv.Itoa(int(config.SConfig.Tcp.Agent.ControlPort)))
// 		log.Fatal(err)
// 	}
// 	logger.Infof("Success listening %s:%d\t%s:%d\n",
// 		config.SConfig.Tcp.Agent.ListenAddr, config.SConfig.Tcp.Agent.ListenPort,
// 		config.SConfig.Tcp.Agent.ControlAddr, config.SConfig.Tcp.Agent.ControlPort)
// 	log.Printf("Begian Listen \n")

// 	// 起协程 使用Listen 返回的Listener类型的连接
// 	// 到这里就完成了监听任务
// 	go connutil.AcConn(ListenLn, heartbeat.HandleClientConn)
// 	connutil.AcConn(CtrlLn, heartbeat.HandleControlConn)

// 	// 阻塞
// }
