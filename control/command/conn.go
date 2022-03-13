package command

import (
	"fmt"
	"net"
	"time"
)

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
	conn.Write([]byte("控制Client,输入操作码\n1. 展示所有连接\t2. 控制连接\t3. 退出\n输入:"))
	cmd, err := conn.Read()
	if err != nil {
		this.conn.Write([]byte("ERR|Failed reading line\r\n"))
		return
	}
	for {

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
