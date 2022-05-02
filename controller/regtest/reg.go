package regtest

import (
	"c2c/controller"
	_ "c2c/tcpagent"
	"fmt"
)

func Reg() {
	fmt.Printf("%v\n", controller.Controller)
	for _, v := range controller.Controller.AgentList {
		fmt.Println(v.GetParams())
	}
}
