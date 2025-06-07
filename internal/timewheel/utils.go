package timewheel

import(
	"timeping/internal/tlog"
	"timeping/pkg/tlist"
)
//确保并发安全以及记录长度，队列中节点过少时触发扩容
func popFront() *tlist.Node {
	//从未使用队列中弹出一个节点
	quemu.Lock()
	defer quemu.Unlock()
	if unuseQueueLen<= 10 {
		tw.expandTaskPool()
	}
	node := UnuseQueue.PopFront()
	if node == nil {
		tlog.Common("unusedqueue is empty", "Warning", "timewheel")
		return nil
	}
	unuseQueueLen--
	return node
}

func pushBack(node *tlist.Node) {
	//放入队尾
	quemu.Lock()
	defer quemu.Unlock()
	unuseQueueLen++
	UnuseQueue.PushBack(node)
}

func delete(t *tlist.Tlist) {
	//删除tlist
	if t == nil {
		tlog.Common("tlist is nil", "Warning", "timewheel")
		return
	}
	quemu.Lock()
	defer quemu.Unlock()

	t.Delete(UnuseQueue)
	unuseQueueLen++
}