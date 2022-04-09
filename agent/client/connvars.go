package globalvars

import (
	"net"
	"sync"
)

// 连接队列
var ConnLi []net.Conn

// 状态队列

type ClientList struct {
	count       int
	Clients     map[int]*Bot
	CliCmdQueue map[int]string
	cntMutex    *sync.Mutex
}

type Bot struct {
	uid     int64
	conn    net.Conn
	CliStat byte
}
