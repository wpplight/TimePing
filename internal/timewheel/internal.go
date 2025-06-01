package timewheel

import (
	"timeping/internal/tlog"
	"timeping/pkg/tlist"
)

func (tw *timewheel)addtask(m chanInfo) {

	if tw.Tw[m.TaskPos].Front() == nil {
		//取未使用节点创建tlist
		//加锁
		tw.TwLocks[m.TaskPos].Lock()
		defer tw.TwLocks[m.TaskPos].Unlock()

		if l, err := tlist.Build(popFront()); err == nil {
			tw.Tw[m.TaskPos].PushBack(&Tasks{
				tl:l,
				timeWheels:m.TaskTime,
			})
			return
		}else{
			tlog.Common("unusedqueue is empty", "Warning", "timewheel")
			//这说明之前没有扩容成功，这里出问题了
			return
		}
	}

	//取出来，修改id
	tp := UnuseQueue.PopFront()
	
	if tp ==nil{
		tlog.Common("unusedqueue is empty", "Warning", "timewheel")
		//列中没有了，出问题了
	}

	for e := tw.Tw[m.TaskPos].Front(); e != nil; e = e.Next() {

		//类型断言
		taskElem, ok := e.Value.(*Tasks)

		//出现错误进行上报然后跳过避免系统崩溃
		if !ok {
			tlog.Common("wrong type", "error", "timewheel")
			tlog.Err_in("wrong type", "error", "timewheel")
			//other 处理
			continue
		}

		//已有轮次，直接插入
		if taskElem.timeWheels == m.TaskTime {
			//加锁
			taskElem.mu.Lock()
			
			//如果tlist不为空，直接插入
			taskElem.tl.PushBack(tp)
			taskElem.mu.Unlock()
			return
		} else if taskElem.timeWheels > m.TaskTime {
			l, err := tlist.Build(popFront())
			if err!=nil{
				//队列中没有了，出问题了
			}
			l.PushBack(tp)
			tw.Tw[m.TaskPos].InsertBefore(&Tasks{
				tl:l,
				timeWheels: m.TaskTime,
			}, e)
			return
		}

	}
	//如果到达尾部，插入
	l, err := tlist.Build(popFront())
	if err!=nil {
		//去扩容
		tlog.Common("unusedqueue is empty", "Warning", "timewheel")
		return
	}
	tw.Tw[m.TaskPos].PushBack(&Tasks{
		tl:l,
		timeWheels: m.TaskTime,
	})
}

//杨神你一定要改啊，这个写的真的一般啊哭哭
func (tw *timewheel)deletetask(m chanInfo) {
	for e := tw.Tw[m.TaskPos].Front(); ; e = e.Next() {
		if taskElem, ok := e.Value.(*Tasks); ok {
			//找到正确位置删除
			if taskElem.timeWheels == m.TaskTime {
				//加锁
				taskElem.mu.Lock()
				defer taskElem.mu.Unlock()
				//找到对应的tlist,现在需要遍历寻找
				for e2 := taskElem.tl.Front(); e2 != nil; e2 = e2.Next {
					if *GetIdPtr(e2) == m.TaskId {
						//找到对应的节点，删除
						e2.Move()
						//将该节点放入未使用队列
						pushBack(e2)
						//如果tlist为空，则删除该轮次的tasks
						if taskElem.tl.IsEmpty() == nil {
							tw.Tw[m.TaskPos].Remove(e)
						}
						break
					}
					if e2.Next == nil {
						tlog.Common("未找到对应的任务，可能已经被删除")
						break
					}
				}

			}
		}
		if e.Next() == nil {
			tlog.Common("未找到对应的任务，可能已经被删除", "Warning", "timewheel")
			break
		}
	}
}
