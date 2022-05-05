package dnsagent

import (
	"c2c/common/def"
	"c2c/controller"
	"c2c/taskqueue"
	"fmt"
)

type Agent struct {
	Params *def.AgentParams

	def.AgentInterface
}

func init() {
	var a Agent = Agent{}
	params, ok := controller.AgentParamsLists["DNSAGENT"]
	if !ok {
		fmt.Println("error")
		return
	}
	a.Params = &params
	controller.RegisterAgent(&a, new(taskqueue.TasksQueue))
}

func (a *Agent) GetParams() def.AgentParams {
	return *a.Params
}
