package utype

type cnf struct{
	Timeinterval uint16//时间轮的间隔
	TaskPoolSize uint16 //任务池的大小
	TimeWheelSize uint16//时间轮的大小
	Port  uint16//使用端口
};
type Node struct{
	Next *Node
	Last *Node
}
type Tlist struct{
	len int
	head *Node
	tail *Node
}
type TimeWheelNode struct {
	Tnode *Node
}
var(
	Conf cnf
	TaskChan = make(chan taskInfo, 1)
)


type taskInfo struct{
	TaskId uint16
	TaskPos uint16
}
