package controller

import (
	"c2c/common/def"
	"c2c/common/logger"
	"c2c/taskqueue"
)

var Controller *Control = NewController()

func NewController() *Control {
	// 从配置中读取地址 端口等信息
	var c *Control = &Control{}
	c.AgentList = make([]def.AgentInterface, 0)
	c.AgentTasksQueue = make(map[string]*taskqueue.TasksQueue)
	return c
}

// 向全局变量添加agent，实现注册机制
func RegisterAgent(i def.AgentInterface, atq *taskqueue.TasksQueue) {
	Controller.AgentList = append(Controller.AgentList, i)
	// 先查询是否已经注册过
	_, ok := Controller.AgentTasksQueue[i.GetParams().Name]
	if !ok {
		logger.Infof("reg a taskqueue to controller")
		Controller.AgentTasksQueue[i.GetParams().Name] = atq
	}
}
