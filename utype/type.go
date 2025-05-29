package utype

type Cnf struct{
	Timeinterval uint16//时间轮的间隔
	TaskPoolSize uint16 //任务池的大小
	TimeWheelSize uint16//时间轮的大小
	Port  uint16//使用端口
};
type Node struct{
	Next *Node
	Last *Node
}
type TimeWheelNode struct {
	Tnode *Node
}

type TaskNode struct {

	Used bool
	TaskId uint16
	Tnode Node
}

type TaskInfo struct{
	TaskId uint16
	TaskPos uint16
}
