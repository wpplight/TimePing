package timewheel

import (
	"container/list"
	"time"
	"timeping/internal/config"
)
var tw *timewheel
// 创建时间轮
func NewTimeWheel() *timewheel {
	tw := new(timewheel)
	tw.taskpool = make([]tasknode, config.Conf.TaskPoolSize)
	tw.Tw = make([]*list.List, config.Conf.TimeWheelSize)
	for i := 0; i < len(tw.Tw); i++ {
		tw.Tw[i] = list.New()
	}
	tw.initialUnuseQueue()
	tw.index = 0
	return tw
}

// 运行时间轮
func (tw *timewheel) Run() {
	tw.ticker = time.NewTicker(time.Duration(config.Conf.Timeinterval) * time.Second)
	defer tw.ticker.Stop()

	//并发处理删改
	go tw.modifyTask()
	//启动计时时间轮
	for {
		<-tw.ticker.C
		go tw.executeTask()
	}

}

func (tw *timewheel) modifyTask() {
	for {
		select {
		case gettask := <-AddTaskChan:

			//增加任务
			tw.addtask(gettask)

		case gettask := <-DeleteTaskChan:
			//删除任务
			tw.deletetask(gettask)
		}
	}
}

func (tw *timewheel) executeTask() {

	
	if tw.index == config.Conf.TimeWheelSize {
		tw.index = 0
		tw.wheel++
	}

	//判断当前节点有没有东西
	if tw.Tw[tw.index].Front() != nil {
		e := tw.Tw[tw.index].Front()
		if taskElem, ok := e.Value.(*Tasks); ok {

			//如果是当前轮次的任务
			if taskElem.timeWheels == tw.wheel {
				//加锁
				taskElem.mu.Lock()
				defer taskElem.mu.Unlock()

				
				tl := taskElem.tl
				//执行该轮次的任务
				
				
				
				for tl.IsEmpty() == nil {
					node := tl.PopFront()
					
					pushBack(node)
				}
				//归还哨兵节点
				delete(tl)
				//清空该轮次的任务
				tw.Tw[tw.index].Remove(e)
			}
		}

	}
	tw.index++
}
