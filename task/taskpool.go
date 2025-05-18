package task

import (
	"container/list"
	t  "timeping/utype"
)

func InitialTaskPool(){
	t.Engine_kernal.TaskPool=list.New();
	TaskPoolNode := make([]t.TaskNode, t.Conf.TaskPoolSize)
	t.Engine_kernal.TaskPool.PushBack(TaskPoolNode)
	creatUnusedTask(TaskPoolNode);
}