package timewheel

import (
	"log"
	"time"
	"timeping/global"
	"timeping/tlist"
	"timeping/task"
)

var Tw []*tlist.Tlist
func InitialTimeWheel() {
	Tw=make([]*tlist.Tlist,global.Conf.TimeWheelSize)
}


func ModifyTask() {
	for{
		select{
			case gettask:=<-global.AddTaskChan:
				//从任务队列中取出任务
				var err error
				if(Tw[gettask.TaskPos]==nil){
					
					Tw[gettask.TaskPos], err = tlist.Build(global.UnuseQueue.PopFront())
					
				}else {
					Tw[gettask.TaskPos].PushFront(global.UnuseQueue.PopBack())
				}
				if err != nil {
					log.Fatal(err)
				}
			case gettask:=<-global.DeleteTaskChan:
				//从任务队列中取出任务
				temp:=Tw[gettask.TaskPos].GetHeadNode()
				for temp!= nil {
					if(task.GetTaskId(temp)==gettask.TaskId){
						
					}
					temp=temp.Next
				}
		}
	}
}
func TimeTicker() {
	ticker := time.NewTicker(time.Duration(global.Conf.Timeinterval) * time.Second)
	defer ticker.Stop()
	var index uint16=0;
	for {
		select {
		case <-ticker.C:

			//判断是否是当前轮的节点，如果是则运行,待写入

			index++;
			if(index==global.Conf.TimeWheelSize){
				index=0;
			}
		}
	}
}
