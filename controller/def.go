package controller

import (
	"c2c/common/def"
	"net"
)

// 定义已有的AGENT TYPE
const (
	TCPAGENT uint8 = 0x1 + iota
	UDPAGENT
	HTTPAGENT
	HTTPSAGENT
	DNSAGENT
	WEBSOCKET
)

// 所有类型的agent
var AgentParamsLists map[string]def.AgentParams = map[string]def.AgentParams{
	"TCPAGENT": {
		Name:      "TCPAGENT",
		AgentType: TCPAGENT,
	},
	"UDPAGENT": {
		Name:      "UDPAGENT",
		AgentType: UDPAGENT,
	},
	"HTTPAGENT": {
		Name:      "HTTPAGENT",
		AgentType: HTTPAGENT,
	},
	"HTTPSAGENT": {
		Name:      "HTTPSAGENT",
		AgentType: HTTPSAGENT,
	},
	"DNSAGENT": {
		Name:      "DNSAGENT",
		AgentType: DNSAGENT,
	},
	"WEBSOCKET": {
		Name:      "WEBSOCKET",
		AgentType: WEBSOCKET,
	},
}

type Control struct {
	BackendListener net.Listener
	AgentList       []def.AgentInterface
}
