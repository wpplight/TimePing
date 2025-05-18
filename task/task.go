package task

import (
	"container/list"
	t "timeping/utype"
)

func creatUnusedTask(TaskPoolNode []t.TaskNode )  {
	t.Engine_kernal.UnuseQueue=list.New()
	for i:=0;i<len(TaskPoolNode);i++{
		TaskPoolNode[i].Used = false
		t.Engine_kernal.UnuseQueue.PushBack(&TaskPoolNode[i].Tnode)
	}
}

