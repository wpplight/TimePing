package task

import (

	"timeping/pkg/tlist"
	"unsafe"
)

func GetTaskId(n *tlist.Node)uint16{
	taskid :=uint16((uintptr(unsafe.Pointer(n))-trans.TaskId))
	return taskid
}

func GetUsed(n *tlist.Node) *bool {
	Used :=(*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(n))-trans.Used))
	return Used
}