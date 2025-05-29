package timewheel

import (
	"log"
	"time"
	"timeping/internal/config"
	"timeping/internal/task"
	"timeping/pkg/tlist"
)

var Tw []*tlist.Tlist
func InitialTimeWheel() {
	Tw=make([]*tlist.Tlist,config.Conf.TimeWheelSize)

}


func ModifyTask() {
	for{
		select{
			case gettask:=<-AddTaskChan:
				//从任务队列中取出任务
				var err error
				if(Tw[aa]==nil){
					Tw[gettask.TaskPos], err = tlist.Build(task.UnuseQueue)
				}else {
					Tw[gettask.TaskPos].PushFront(task.UnuseQueue.PopBack())
				}
				if err != nil {
					log.Fatal(err)
				}
			case gettask:=<-DeleteTaskChan:
				//从任务队列中取出任务
				temp:=Tw[gettask.TaskPos].PopFront()
				for temp!= nil {
					if(task.GetTaskId(temp)==gettask.TaskId){
						
					}
					temp=temp.Next
				}
		}
	}
}
func TimeTicker() {
	ticker := time.NewTicker(time.Duration(config.Conf.Timeinterval) * time.Second)
	defer ticker.Stop()
	var index uint16=0;
	for {
		select {
		case <-ticker.C:

			//判断是否是当前轮的节点，如果是则运行,待写入

			index++;
			if(index==config.Conf.TimeWheelSize){
				index=0;
			}
		}
	}
}
