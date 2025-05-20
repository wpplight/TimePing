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
	l:=u.TaskNode{}
	global.Trans.TCB =unsafe.Offsetof(l.Tnode)
}
func creatUnusedTask(TaskPoolNode []u.TaskNode) {
	for i:=0;i<len(TaskPoolNode);i++{
		TaskPoolNode[i].Used = false
		TaskPoolNode[i].TaskId  = uint16(i)
		global.UnuseQueue.PushBack(&TaskPoolNode[i].Tnode)
	}
}
// 将Node转换为TaskNode
func N2Tcb(n *u.Node) *u.TaskNode {
	return (*u.TaskNode)(unsafe.Pointer(uintptr(unsafe.Pointer(n))-global.Trans.TCB))
}

