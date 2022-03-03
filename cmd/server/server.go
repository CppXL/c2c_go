package main

import (
	"c2c/common/utils/connutil"
	"c2c/config"
	"fmt"
	"log"
	"net"
	"runtime"
	"strconv"
)

// 监听端口
// 把已有连接加入到队列内 维护连接状态

var (
	// 监听端口和控制端口连接
	ListenLn, CtrlLn net.Listener
)

func init() {
	// 解析命令行参数
	var err error
	err = config.InitSrvConfig("")
	if err != nil {
		log.Fatal(err)
	}
	// 没错误就绑定端口
	ListenLn, err = net.Listen("tcp", config.SConfig.Server.ListenAddr+":"+strconv.Itoa(int(config.SConfig.Server.ListenPort)))
	if err != nil {
		fmt.Println(config.SConfig.Server.ListenAddr + strconv.Itoa(int(config.SConfig.Server.ListenPort)))
		log.Fatal(err)
	}
	CtrlLn, err = net.Listen("tcp", config.SConfig.Server.ControlAddr+":"+strconv.Itoa(int(config.SConfig.Server.ControlPort)))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Success listening %s:%d\t%s:%d\n",
		config.SConfig.Server.ListenAddr, config.SConfig.Server.ListenPort,
		config.SConfig.Server.ControlAddr, config.SConfig.Server.ControlPort)

}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf("Begian Listen \n")

	go connutil.AcConn(ListenLn, connutil.HandleClientConn)
	connutil.AcConn(CtrlLn, connutil.HandleControlConn)

	// 阻塞
}
