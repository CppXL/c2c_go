package def

// 定义agent类型
const (
	TCPAGENT uint8 = 0x1 + iota
	UDPAGENT
	HTTPAGENT
	HTTPSAGENT
	DNSAGENT
	WEBSOCKET
)
