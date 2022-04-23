package controller

type agenttype uint8

const (
	TCPAGENT agenttype = 0x1 + iota
	HTTPAGENT
)
