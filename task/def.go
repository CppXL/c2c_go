package task

type TaskList struct {
	AgentType   uint8
	ClientToken uint32
	TaskID      uint32
	Cmd         uint16
	CmdLen      uint16
	CmdArgs     []byte
}

func NewTaskList(agentType uint8) *TaskList {
	t := &TaskList{AgentType: agentType}

	return t

}

func (t *TaskList) Watch() {

}
