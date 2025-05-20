package task

import (
	"timeping/utype"
	"timeping/global"
)
type TaskNode struct {
	taskId uint16
	Used bool 
	Tnode utype.Node
};
//返回一个总的数据池，和一个初始化的持有多任务块的未使用任务队列
func InitialTaskPool() {
	global.Taskpool.PushBack(make([]TaskNode,utype.Conf.TaskPoolSize))
	creatUnusedTask(global.Taskpool.Front().Value.([]TaskNode))
}

