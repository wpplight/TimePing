package utype
import "container/list"

type cnf struct{
	Timeinterval uint16//时间轮的间隔
	TaskPoolSize uint16 //任务池的大小
	TimeWheelSize uint16//时间轮的大小
	Port  uint16//使用端口
};
type Node struct{
	next *Node
	last *Node
}
type TaskNode struct {
	taskId uint16
	Used bool 
	Tnode Node
};
var Engine_kernal struct{
	TaskPool *list.List
	UnuseQueue *list.List
}
type TimeWheelNode struct{
	node *TaskNode
}

var(
	Conf cnf
)
type taskInfo struct{
	TaskId uint16
	TaskPos uint16
}
var(
	TaskChan = make(chan taskInfo, 1)
)