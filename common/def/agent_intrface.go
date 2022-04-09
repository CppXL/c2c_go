package def

import "net"

// agent function interface def
type AgentInterface interface {
	SendCmd(conn net.Conn, cmd uint16, cmdArgs []byte) error
	OnHeartBeatTimeout() // heartbeat time out
	OnClientClose()      // client close
}
