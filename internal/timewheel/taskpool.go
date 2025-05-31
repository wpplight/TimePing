package timewheel

import(
	"timeping/pkg/tlist"
	"unsafe"
)

type tasknode struct{
	node *tlist.Node
	id int
}
// taskpool是一个任务池，预分配1000的容量
var taskpool = make([]*tasknode,1000)


// UnuseQueue是一个未使用的任务队列
var UnuseQueue = tlist.New()
func InitialUnuseQueue() {
	//将taskpool中的所有节点放入UnuseQueue中
	for i := 0; i < len(taskpool); i++ {
		UnuseQueue.PushBack(taskpool[i].node)
	}
}


func GetIdPtr(node *tlist.Node) *int {
    // 将node指针转换为tasknode指针，并返回id字段的指针
    taskPtr := (*tasknode)(unsafe.Pointer(
        uintptr(unsafe.Pointer(node)) - unsafe.Offsetof(tasknode{}.node),
    ))
    return &taskPtr.id
}