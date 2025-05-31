package timewheel

import(
	"timeping/pkg/tlist"
	"unsafe"
)






func (tw *timewheel) initialUnuseQueue() {
	spos=unsafe.Offsetof(tasknode{}.node)
	//将taskpool中的所有节点放入UnuseQueue中
	for i := 0; i < len(tw.taskpool); i++ {
		UnuseQueue.PushBack(&tw.taskpool[i].node)
	}
}

//将Node转换成tasknode
func GetTask(node *tlist.Node) *tasknode{
	return  (*tasknode)(unsafe.Pointer(uintptr(unsafe.Pointer(node)) -spos))
}

func GetIdPtr(node *tlist.Node) *int {
	return &GetTask(node).id
}

