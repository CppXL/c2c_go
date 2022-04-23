package connutil

import (
	"c2c/common/logger"
	"net"
)

func AcConn(ln net.Listener, handleConnFunc func(conn net.Conn, args ...interface{}), args ...interface{}) {
	for {
		// 接受连接，没有连接这里会阻塞
		conn, err := ln.Accept()
		if err != nil {
			// 打印错误 继续循环
			logger.Warnf(err.Error())
			continue
		}
		logger.Infof("Accept Conn from %v to %v  handle by %v\n", conn.RemoteAddr(), conn.LocalAddr(), handleConnFunc)
		go handleConnFunc(conn, args)
	}
}
