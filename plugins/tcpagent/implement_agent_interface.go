package tcpagent

import (
	"c2c/common/def"
	"fmt"
)

func (a *Agent) Shell() error {

	fmt.Println(a)
	return nil
}
func (a *Agent) GetParams() def.AgentParams {
	return *a.Params
}
