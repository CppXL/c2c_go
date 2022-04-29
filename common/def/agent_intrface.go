package def

import (
	"net"
)

// 所有类型的agent都要实现这个接口
type AgentInterface interface {
	// 进入agent的shell
	Shell() error
	// 关闭agent
	Close() error
}

type Agent struct {
	FrontListener net.Listener
}
