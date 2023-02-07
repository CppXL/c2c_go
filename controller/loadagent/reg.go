package loadagent

import (
	"c2c/controller"
	_ "c2c/plugins/dnsagent"
	_ "c2c/plugins/tcpagent"
	"fmt"
)

func RegAgent() {
	fmt.Printf("%+v\n", controller.Controller)
	for _, v := range controller.Controller.AgentList {
		fmt.Println(v.GetParams())
	}

	// task := controller.Controller.AgentTasksQueue["DNSAGENT"]
	// task.TaskQueue = append(task.TaskQueue, new(taskqueue.Task))
	// task.TaskQueue = append(task.TaskQueue, new(taskqueue.Task))
	// task.TaskQueue[0].ClientToken = 0x355
	// for _, v := range controller.Controller.AgentTasksQueue {
	// 	for _, g := range v.TaskQueue {
	// 		fmt.Printf("%T\n", *&g.ClientToken)
	// 	}
	// }
}
