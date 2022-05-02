package def

// 所有类型的agent都要实现这个接口，agent通过register机制注册到controller中
type AgentInterface interface {
	// 进入agent的shell
	Shell() error
	// 关闭agent
	Close() error
	// 重启agent
	Restart()
	// 停止agent
	Stop()
	// 返回agent相关的参数
	GetParams() AgentParams
}

// agnet的一些基本信息
type AgentParams struct {
	Name      string
	AgentType uint8
}
