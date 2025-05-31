package timewheel

import (
	"timeping/internal/tlog"
	"timeping/pkg/tlist"
)

func (tw *timewheel)addtask(m chanInfo) {

	if tw.Tw[m.TaskPos].Front() == nil {
		//取未使用节点创建tlist
		if l, err := tlist.Build(UnuseQueue.PopFront()); err == nil {
			tw.Tw[m.TaskPos].PushBack(&Tasks{
				l,
				m.TaskTime,
			})
		}else{
			//去扩容
		}
	}

	//取出来，修改id
	tp := UnuseQueue.PopFront()
	if tp ==nil{
		tlog.Common("unusedqueue is empty", "Warning", "timewheel")
		//调用扩容函数等待扩容
	}
	innode := GetTask(tp)
	innode.id = m.TaskId

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
			taskElem.tl.PushBack(tp)
			return
		} else if taskElem.timeWheels > m.TaskTime {
			l, err := tlist.Build(UnuseQueue.PopFront())
			if err!=nil{
				//去扩容
			}
			l.PushBack(tp)
			tw.Tw[m.TaskPos].InsertBefore(&Tasks{l, m.TaskTime}, e)
			return
		}

	}
	//如果到达尾部，插入
	l, err := tlist.Build(UnuseQueue.PopFront())
	if err!=nil {
		//去扩容
	}
	tw.Tw[m.TaskPos].PushBack(&Tasks{
		l,
		m.TaskTime,
	})
}

//杨神你一定要改啊，这个写的真的一般啊哭哭
func (tw *timewheel)deletetask(m chanInfo) {
	for e := tw.Tw[m.TaskPos].Front(); e != nil; e = e.Next() {
		if taskElem, ok := e.Value.(*Tasks); ok {
			//找到正确位置删除
			if taskElem.timeWheels == m.TaskTime {
				//找到对应的tlist,现在需要遍历寻找
				for e2 := taskElem.tl.Front(); e2 != nil; e2 = e2.Next {
					if *GetIdPtr(e2) == m.TaskId {
						//找到对应的节点，删除
						e2.Move()
						//将该节点放入未使用队列
						UnuseQueue.PushBack(e2)
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
	}
}
