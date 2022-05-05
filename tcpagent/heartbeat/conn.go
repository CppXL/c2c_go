package heartbeat

import (
	"bufio"
	"c2c/common/logger"
	"io"
	"net"
)

//
func HandleClientConn(conn net.Conn, args ...interface{}) {
	defer func() {
		conn.Close()
	}()

}

// connect to control
func HandleControlConn(conn net.Conn, args ...interface{}) {
	defer func() {
		conn.Close()
	}()

	// conn.SetDeadline(time.Now().Add(10 * time.Second))
	reader := bufio.NewReader(conn)
	// var in string
	var i int
	for {
		conn.Write([]byte("console>"))
		cmd, err := reader.ReadString('\n')
		if err == io.EOF {
			logger.Infof("recv EOF")
			return
		} else {
			logger.FatalIfError(err)

		}
		logger.Infof("recv %s form %v\n", cmd, conn.RemoteAddr())
		// fmt.Scanln(&in)

		if cmd[len(cmd)-1] == '\n' {
			conn.Write([]byte(cmd))

		} else {
			conn.Write([]byte(cmd + "\n"))

		}
		i++
		if i == 10 {
			break
		}
	}

}
func StoreCliConn(conn net.Conn, connLi []net.Conn) []net.Conn {
	connLi = append(connLi, conn)
	return connLi
}
