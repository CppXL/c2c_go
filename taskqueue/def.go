package task

// 任务队列结构体
type TaskList struct {
	AgentType   uint8
	ClientToken uint32
	TaskID      uint32
	TaskStatus  int8
	Cmd         uint16
	CmdLen      uint16
	CmdArgs     []byte
}

// 任务状态标志
const (
	TASKFINISH = 1 + iota
	TASKFAILED
	TASKEXECING
	TASKTIMEOUT
)

func NewTaskList(agentType uint8) *TaskList {
	t := &TaskList{AgentType: agentType}
	return t
}

func (t *TaskList) Watch() {

}
