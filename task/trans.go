package task
import (
	"unsafe"
	"timeping/global"
	"timeping/utype"
)

func GetTaskId(n *utype.Node)uint16{
	taskid :=uint16((uintptr(unsafe.Pointer(n))-global.Trans.TaskId))
	return taskid
}

func GetUsed(n *utype.Node) *bool {
	Used :=(*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(n))-global.Trans.Used))
	return Used
}