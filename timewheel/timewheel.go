package timewheel

import (
	"timeping/utype"
	
)

var Tw []utype.TimeWheelNode
func InitialTimeWheel() {
	Tw=make([]utype.TimeWheelNode,utype.Conf.TimeWheelSize)
}

func AddTask() {
	index:=0 //获取逻辑待实现
	for{
		select{
			case task:=<-utype.TaskChan:
				//从任务队列中取出任务
				Tw[task.TaskPos]=&utype.Engine_kernal.UnuseQueue.Front()
				
		}
	}
}
