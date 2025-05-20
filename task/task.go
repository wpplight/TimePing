package task

import (
	"timeping/global"
)

//通过创建好的任务池，创建一个未使用的任务队列
func creatUnusedTask(TaskPoolNode []TaskNode) {
	for i:=0;i<len(TaskPoolNode);i++{
		TaskPoolNode[i].Used = false
		global.UnuseQueue.PushBack(&TaskPoolNode[i].Tnode)
	}
}
