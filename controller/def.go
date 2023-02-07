package controller

import (
	"c2c/common/def"
	"c2c/taskqueue"
	"net"
)

// 定义已有的AGENT TYPE

// 所有类型的agent
var AgentParamsLists map[string]def.AgentParams = map[string]def.AgentParams{
	"TCPAGENT": {
		Name:      "TCPAGENT",
		AgentType: def.TCPAGENT,
		Status:    def.AgentStatus_NotLoad,
	},
	"UDPAGENT": {
		Name:      "UDPAGENT",
		AgentType: def.UDPAGENT,
		Status:    def.AgentStatus_NotLoad,
	},
	"HTTPAGENT": {
		Name:      "HTTPAGENT",
		AgentType: def.HTTPAGENT,
		Status:    def.AgentStatus_NotLoad,
	},
	"HTTPSAGENT": {
		Name:      "HTTPSAGENT",
		AgentType: def.HTTPSAGENT,
		Status:    def.AgentStatus_NotLoad,
	},
	"DNSAGENT": {
		Name:      "DNSAGENT",
		AgentType: def.DNSAGENT,
		Status:    def.AgentStatus_NotLoad,
	},
	"WEBSOCKET": {
		Name:      "WEBSOCKET",
		AgentType: def.WEBSOCKET,
		Status:    def.AgentStatus_NotLoad,
	},
}

type Control struct {
	BackendListener net.Listener                     // 提供给后端连接的连接
	AgentList       []def.AgentInterface             // 实现了AgentInterface接口的agent列表
	AgentTasksQueue map[string]*taskqueue.TasksQueue //这里要指针，因为需要每次取值的时候取到的是同一个struct
}
