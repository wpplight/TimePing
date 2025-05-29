package task

import (
	"container/list"
	"timeping/pkg/tlist"
)


type TaskNode struct {
	Used bool
	TaskId uint16
	Tnode tlist.Node
}

var(
	taskpool =list.New()
	unuseQueue =tlist.New()

	trans struct{
	TaskId uintptr
	Used uintptr
	}
)