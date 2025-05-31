package timewheel

import (
	"container/list"
	"time"
	"timeping/internal/config"
)

//创建时间轮
func NewTimeWheel() *timewheel {
	tw:=new(timewheel)
	tw.Tw =make([]*list.List,config.Conf.TimeWheelSize)
	for i:=0;i<len(tw.Tw);i++{
		tw.Tw[i]=list.New()
	}
	tw.initialUnuseQueue()
	tw.index=0
	return tw
}
//运行时间轮
func(tw *timewheel) Run(){
	tw.ticker=time.NewTicker(time.Duration(config.Conf.Timeinterval) * time.Second)
	defer tw.ticker.Stop()

	//并发处理删改
	go tw.modifyTask()
	//启动计时时间轮
	for {
		<-tw.ticker.C
		go tw.executeTask()
	}

}

func (tw *timewheel)modifyTask() {
	for{
		select{
			case gettask:=<-AddTaskChan:

				//增加任务
				tw.addtask(gettask)

			case gettask:=<-DeleteTaskChan:
				//删除任务
				tw.deletetask(gettask)
		}
	}
}

func (tw *timewheel) executeTask() {

	wheel :=0 //轮次
		tw.index++
		if(tw.index==config.Conf.TimeWheelSize){
			tw.index=0;
		}

		//判断当前节点有没有东西
		if tw.Tw[tw.index].Front() != nil {
				e:=tw.Tw[tw.index].Front()
				if taskElem, ok := e.Value.(*Tasks); ok {
					
					//如果是当前轮次的任务
					if taskElem.timeWheels == wheel {
						tl:=taskElem.tl
						//执行该轮次的任务
						for tl.IsEmpty()==nil {
							node:=tl.PopFront()
							//处理函数
							UnuseQueue.PushBack(node)
						}
						//归还哨兵节点
						tl.Delete(UnuseQueue)
						//清空该轮次的任务
						tw.Tw[tw.index].Remove(e)
					}
				}
			
			
		}
		
	
}

