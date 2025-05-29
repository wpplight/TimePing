package task

import (
	"timeping/internal/config"
	"unsafe"
)

//返回一个总的数据池，和一个初始化的持有多任务块的未使用任务队列
func InitialTaskPool() {

	taskpool.PushBack(make([]TaskNode,config.Conf.TaskPoolSize))

	creatUnusedTask(taskpool.Front().Value.([]TaskNode))
 
	var t =TaskNode{}
	trans.TaskId = unsafe.Offsetof(t.Tnode)-unsafe.Offsetof(t.TaskId)
	trans.Used = unsafe.Offsetof(t.Tnode)-unsafe.Offsetof(t.Used)
	
}

func creatUnusedTask(taskPoolNode []TaskNode) {

	for i:=0;i<len(taskPoolNode);i++{

		taskPoolNode[i].Used = false

		taskPoolNode[i].TaskId  = uint16(i)

		unuseQueue.PushBack(&taskPoolNode[i].Tnode)
	}

}