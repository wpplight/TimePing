package task

import (
	"container/list"
	"timeping/tlist"
	"timeping/utype"
)
type TaskNode struct {
	taskId uint16
	Used bool 
	Tnode utype.Node
};
//返回一个总的数据池，和一个初始化的持有多任务块的未使用任务队列
func InitialTaskPool() (*list.List,*tlist.Tlist) {
	taskpool := list.New()
	taskpool.PushBack(make([]TaskNode,utype.Conf.TaskPoolSize))
	return taskpool,creatUnusedTask(taskpool.Front().Value.([]TaskNode))
}

