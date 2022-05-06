package def

import "c2c/taskqueue"

// 所有类型的agent都要实现这个接口，agent通过register机制注册到controller中

type AgentInterface interface {
	Shell() error                            // 进入agent的shell
	Close() error                            // 关闭agent
	Restart()                                // 重启agent
	Stop()                                   // 停止agent
	GetParams() AgentParams                  // 返回agent相关的参数
	OnTaskAdded(task *taskqueue.Task) uint32 // 当任务被添加到agent中时调用
	OnTaskChange(task *taskqueue.Task)       // 当任务状态发生变化时调用
	OnClientConnected() uint32               // 当client连接到agent时调用，返回一个token
	OnClientDisconnected()                   // 当client断开连接时调用
}

// agnet的一些基本信息
type AgentParams struct {
	Name      string
	AgentType uint8
	Status    AgentStatus
}

type AgentStatus uint8

// agent status def
const (
	AgentStatus_NotLoad AgentStatus = iota + 1
	AgentStatus_LoadFailed
	AgentStatus_Loaded
	AgentStatus_Running
	AgentStatus_Runfailed
	AgentStatus_Stopped
	AgentStatus_Restarting
)

func (a AgentStatus) String() string {
	switch a {
	case AgentStatus_NotLoad:
		return "AgentStatus_NotLoad"
	case AgentStatus_LoadFailed:
		return "AgentStatus_LoadFailed"
	case AgentStatus_Running:
		return "AgentStatus_Running"
	case AgentStatus_Stopped:
		return "AgentStatus_Stopped"
	case AgentStatus_Restarting:
		return "AgentStatus_Restarting"
	default:
		return "AgentStatus_Unknown"
	}
}
