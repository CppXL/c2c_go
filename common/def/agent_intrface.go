package def

import (
	"net"
)

// 所有类型的agent都要实现这个接口
type AgentInterface interface {
	SendCmd(conn net.Conn, cmd uint16, cmdArgs []byte) error
	OnHeartBeatTimeout() // heartbeat time out
	OnClientClose()      // client close
}

type Agent struct {
	FrontListener net.Listener
}
