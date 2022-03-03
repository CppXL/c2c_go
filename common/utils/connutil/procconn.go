package connutil

import (
	"fmt"
	"log"
	"net"
	"time"
)

func AcConn(ln net.Listener, handleConn func(conn net.Conn)) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		log.Printf("Accept Conn from %v to %v \n", conn.RemoteAddr(), conn.LocalAddr())
		// 存储连接

		// 处理和clint的连接
		go handleConn(conn)

	}
}

//
func HandleClientConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

}

// 控制台连接
func HandleControlConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	conn.SetDeadline(time.Now().Add(10 * time.Second))
	for {
		var code int
		fmt.Print("控制Client,输入操作码\n1. 展示所有连接\t2. 控制连接\t3. 退出\n输入:")
		fmt.Scan(&code)
		switch code {
		case 1:
			// contrcli.ShowAllCli(globalvars.ConnLi)
			continue
		case 2:
			// contrcli.ShowAllCli(globalvars.ConnLi)
			var s int
			fmt.Print("输入控制那个client:")
			fmt.Scan(&s)
			// contrcli.Snd2Cli(globalvars.ConnLi[s])
			continue
		}
	}

}
func StoreCliConn(conn net.Conn, connLi []net.Conn) []net.Conn {
	connLi = append(connLi, conn)
	return connLi
}
