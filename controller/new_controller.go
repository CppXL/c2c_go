package controller

import "c2c/common/def"

var Controller *Control = NewController()

func NewController() *Control {
	return &Control{}
}

// 向全局变量添加agent，实现注册机制
func RegisterAgent(i def.AgentInterface) {
	Controller.AgentList = append(Controller.AgentList, i)
}
