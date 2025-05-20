package task

import (
	"timeping/global"
	u "timeping/utype"
	"unsafe"
)

//返回一个总的数据池，和一个初始化的持有多任务块的未使用任务队列
func InitialTaskPool() {
	global.Taskpool.PushBack(make([]u.TaskNode,global.Conf.TaskPoolSize))
	creatUnusedTask(global.Taskpool.Front().Value.([]u.TaskNode))
	
}
func creatUnusedTask(TaskPoolNode []u.TaskNode) {
	for i:=0;i<len(TaskPoolNode);i++{
		TaskPoolNode[i].Used = false
		global.UnuseQueue.PushBack(&TaskPoolNode[i].Tnode)
	}
}

func N2Tcb(n *u.Node) (*u.TaskNode,error) {
	unsafe

}

