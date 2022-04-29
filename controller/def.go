package controller

import "net"

type agenttype uint8

const (
	TCPAGENT agenttype = 0x1 + iota
	HTTPAGENT
)

type Controller struct {
	BackendListener net.Listener
}
