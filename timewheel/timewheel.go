package timewheel

import (
	"time"
	"timeping/global"
	"timeping/utype"
	"timeping/tlist"
)

var Tw []*tlist.Tlist
func InitialTimeWheel() {
	Tw=make([]*tlist.Tlist,global.Conf.TimeWheelSize)
}


func AddTask() {
	for{
		select{
			case task:=<-global.TaskChan:
				//从任务队列中取出任务
				if(Tw[task.TaskPos].Tnode==nil){
					Tw[task.TaskPos].Tnode=global.UnuseQueue.PopFront()
				}else {
					for Tw[task.TaskPos].Tnode != nil {
						Tw[task.TaskPos].Tnode=Tw[task.TaskPos].Tnode.Next
					}
					Tw[task.TaskPos].Tnode=global.UnuseQueue.PopFront()
				}
				
		}
	}
}
func TimeTicker() {
	ticker := time.NewTicker(time.Duration(utype.Conf.Timeinterval) * time.Second)
	defer ticker.Stop()
	var index uint16=0;
	for {
		select {
		case <-ticker.C:
			//判断是否是当前轮的节点，如果是则运行,待写入
			index++;
			if(index==utype.Conf.TimeWheelSize){
				index=0;
			}
		}
	}
}
