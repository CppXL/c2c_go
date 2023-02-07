package taskqueue

// 任务队列结构体
type Task struct {
	ClientToken uint32
	TaskID      uint32
	TaskStatus  int8
	Cmd         uint16
	CmdLen      uint16
	CmdArgs     []byte
}

// 任务队列定义
type TasksQueue struct {
	TaskQueue []*Task
}

// 任务状态标志
const (
	TASKFINISH uint8 = 1 + iota
	TASKFAILED
	TASKEXECING
	TASKTIMEOUT
)
