package timewheel

import (
	"container/list"
	"sync"
	"time"
	"timeping/pkg/tlist"
)



type chanInfo struct {
	TaskId   int
	TaskPos  int
	TaskTime int
}
type tasknode struct{
	id int
	node tlist.Node
}
type Tasks struct {
	tl *tlist.Tlist
	timeWheels int
	mu sync.Mutex
}
type timewheel struct{
	
	taskpool []tasknode //内存池
	Tw [] *list.List//时间轮结构
	ticker *time.Ticker//计时器
	TwLocks  []sync.Mutex //时间轮锁
	index uint16//时间轮当前下标
	wheel int //时间轮当前轮次
}

var (
	AddTaskChan    = make(chan chanInfo, 1)
	DeleteTaskChan = make(chan chanInfo, 1)
	UpdateTaskChan = make(chan chanInfo, 1)
	
	spos uintptr //偏移

	// UnuseQueue是一个未使用的任务队列
	UnuseQueue = tlist.New()
)
