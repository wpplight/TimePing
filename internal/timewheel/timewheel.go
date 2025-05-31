package timewheel

import (
	"container/list"
	"log"
	"time"
	"timeping/internal/config"
	"timeping/internal/task"
	"timeping/pkg/tlist"
)
var Tw []*list.List
func InitialTimeWheel() {
	Tw =make([]*list.List,config.Conf.TimeWheelSize)
}
type Tasks struct {
	tl *tlist.Tlist
	timeWheels int
}
func ModifyTask() {
	for{
		select{
			case gettask:=<-AddTaskChan:
				//增加任务

				id :=GEtId(UnuseQueue.Front())


				if Tw[gettask.TaskPos].Front()==nil{
					Tw[gettask.TaskPos] = list.New()
					Tw[gettask.TaskPos].PushBack(&Tasks{
						tlist.New(),
						gettask.TaskTime,
					})
				}
				for e := Tw[gettask.TaskPos].Front(); ; e = e.Next() {
					//找到正确位置插入
					if taskElem, ok := e.Value.(*Tasks); ok {

						if taskElem.timeWheels == gettask.TaskTime {
							taskElem.tl.PushBack(UnuseQueue.PopFront())
						}else if taskElem.timeWheels > gettask.TaskTime {
							Tw[gettask.TaskPos].InsertBefore(&Tasks{
								tlist.New(),
								gettask.TaskTime,
							}
						}

					}
					if e == nil {
						//如果到达尾部，插入
						Tw[gettask.TaskPos].PushBack(&Tasks{
							tlist.New(),
							gettask.TaskTime,
						})
						break
					}
				}

				

			case gettask:=<-DeleteTaskChan:
				//从任务队列中取出任务
				
		}
	}
}


func TimeTicker() {
	ticker := time.NewTicker(time.Duration(config.Conf.Timeinterval) * time.Second)
	defer ticker.Stop()
	var index uint16=0;
	for {
		<-ticker.C

			//判断是否是当前轮的节点，如果是则运行,待写入

			index++;
			if(index==config.Conf.TimeWheelSize){
				index=0;
			}
		
	}
}
