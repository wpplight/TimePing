package timewheel

import (
	"sync"
	"timeping/pkg/tlist"
	"unsafe"
)



var unuseQueueLen int //未使用队列长度
var quemu sync.Mutex //互斥锁

func (tw *timewheel) initialUnuseQueue(taskpool ...[]tasknode) {
	//如果传入了taskpool则使用传入的taskpool，否则使用当前的taskpool
	spos=unsafe.Offsetof(tasknode{}.node)
	//将taskpool中的所有节点放入UnuseQueue中
	for i := 0; i < len(tw.taskpool); i++ {
		UnuseQueue.PushBack(&tw.taskpool[i].node)
	}
	unuseQueueLen = len(tw.taskpool) //更新未使用队列长度
}

//将Node转换成tasknode
func GetTask(node *tlist.Node) *tasknode{
	return  (*tasknode)(unsafe.Pointer(uintptr(unsafe.Pointer(node)) -spos))
}

func GetIdPtr(node *tlist.Node) *int {
	return &GetTask(node).id
}

//扩容任务池
func (tw *timewheel)expandTaskPool() {
	//扩容任务池
	newpool := make([]tasknode, len(NewTimeWheel().taskpool)/4+1)
	tw.initialUnuseQueue(newpool) //重新初始化未使用队列
	NewTimeWheel().taskpool = append(tw.taskpool, newpool...)
	
}