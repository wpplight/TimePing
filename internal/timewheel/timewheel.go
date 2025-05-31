package timewheel

import (
	"container/list"
	"log"
	"time"
	"timeping/internal/config"
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

				*GetIdPtr(UnuseQueue.Front())= gettask.TaskId


				if Tw[gettask.TaskPos].Front()==nil{
					Tw[gettask.TaskPos] = list.New()
					l,_:= tlist.Build(UnuseQueue.PopFront())
					Tw[gettask.TaskPos].PushBack(&Tasks{
						l,
						gettask.TaskTime,
					})
				}
				for e := Tw[gettask.TaskPos].Front(); ; e = e.Next() {
					//找到正确位置插入
					if taskElem, ok := e.Value.(*Tasks); ok {
						//已有轮次，直接插入
						if taskElem.timeWheels == gettask.TaskTime {
							taskElem.tl.PushBack(UnuseQueue.PopFront())
							break
						}else if taskElem.timeWheels > gettask.TaskTime {
							l,_:= tlist.Build(UnuseQueue.PopFront())
							Tw[gettask.TaskPos].InsertBefore(&Tasks{
								l,
								gettask.TaskTime,
							},e)
							break
						}

					}
					if e == nil {
						//如果到达尾部，插入
						l,_:= tlist.Build(UnuseQueue.PopFront())
						Tw[gettask.TaskPos].PushBack(&Tasks{
							l,
							gettask.TaskTime,
						})
						break
					}
				}

				

			case gettask:=<-DeleteTaskChan:
				//删除任务
				for e := Tw[gettask.TaskPos].Front(); e != nil; e = e.Next() {
					if taskElem, ok := e.Value.(*Tasks); ok {
						//找到正确位置删除
						if taskElem.timeWheels == gettask.TaskTime {
							//找到对应的tlist,现在需要遍历寻找
							for e2 := taskElem.tl.Front(); e2 != nil; e2 = e2.Next {
								if *GetIdPtr(e2) == gettask.TaskId {
									//找到对应的节点，删除
									e2.Move()
									//将该节点放入未使用队列
									UnuseQueue.PushBack(e2)
									//如果tlist为空，则删除该轮次的tasks
									if taskElem.tl.IsEmpty()==nil {
										Tw[gettask.TaskPos].Remove(e)
									}
									break
								}
								if e2.Next==nil {
									log.Println("未找到对应的任务，可能已经被删除")
									break
								}
							}

						}
					}
				}
				
		}
	}
}
func ExecuteTask(id int) {
}

func TimeTicker() {
	ticker := time.NewTicker(time.Duration(config.Conf.Timeinterval) * time.Second)
	defer ticker.Stop()
	var index uint16=0;
	wheel :=0 //轮次
	for {
		<-ticker.C

			//判断是否是当前轮的节点，如果是则运行
		if Tw[index] != nil && Tw[index].Front() != nil {
			for e := Tw[index].Front(); e != nil; e = e.Next() {
				if taskElem, ok := e.Value.(*Tasks); ok {
					//如果是当前轮次的任务
					if taskElem.timeWheels == wheel {
						//执行该轮次的任务
						for e2 := taskElem.tl.Front(); e2 != nil; e2 = e2.Next {
							ExecuteTask(*GetIdPtr(e2))
							//将该节点放入未使用队列
							UnuseQueue.PushBack(e2)
						}
						//清空该轮次的任务
						if taskElem.tl.IsEmpty()==nil {
							Tw[index].Remove(e)
						}
					}
				}
			}
			
		}
			
			
		index++;
		if(index==config.Conf.TimeWheelSize){
			index=0;
			wheel++;
			//wheel重置功能待写入
		}
		
	}
}
