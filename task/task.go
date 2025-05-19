package task

import (
	"timeping/tlist"
)

//通过创建好的任务池，创建一个未使用的任务队列
func creatUnusedTask(TaskPoolNode []TaskNode) *tlist.Tlist {
	UnuseQueue:=tlist.New()
	for i:=0;i<len(TaskPoolNode);i++{
		TaskPoolNode[i].Used = false
		UnuseQueue.PushBack(&TaskPoolNode[i].Tnode)
	}
	return UnuseQueue
}
