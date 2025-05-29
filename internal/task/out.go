package task

import (
	"timeping/internal/config"
	"timeping/pkg/tlist"
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

func GetNode() (*tlist.Node,error){
	if UnuseQueue.IsEmpty() !=nil{
		return nil,UnuseQueue.IsEmpty()
	}
	return UnuseQueue.PopFront(),nil
}