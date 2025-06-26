package tlist

import (
	"fmt"
	"timeping/internal/tlog"
)

// Tlist 是一个侵入式双向循环链表结构
// 使用哨兵节点(root)作为链表边界标识
type Tlist struct {
	root   *Node  // 哨兵节点，既是头也是尾
}

// init 初始化链表
// 创建一个哨兵节点并设置其Last和Next都指向自己
// 形成空链表的初始状态
func (t *Tlist) init() *Tlist {
	t.root = new(Node)
	t.root.Last = t.root // 指向队尾
	t.root.Next = t.root // 指向队首
	return t
}

// Build 使用已有内存块创建一个链表
// 参数n必须非空，否则返回错误
// 常用于复用已有节点重建链表
func Build(n *Node) (*Tlist, error) {
	if n == nil {
		tlog.Common("Tlist build error nil", "tlist")
		tlog.Err_in("Tlist build, pointer is nil", "tlist")
		return nil, fmt.Errorf("Tlist init error")
	}
	k := new(Tlist)
	
	k.root = n

	return k, nil
}

// IsEmpty 检查链表是否为空
// 空链表是指哨兵节点的Next指向自己
// 如果链表为空返回error("empty")
func (n *Tlist) IsEmpty() error {
	if n == nil {
		tlog.Common("空指针被使用in empty check", "tlist")
		tlog.Err_in("空指针被使用in empty check", "tlist")
		return fmt.Errorf("in empty check")
	}
	if n.root.Next == n.root {
		return fmt.Errorf("empty")
	}
	return nil
}

// New 创建一个新的侵入式链表
// 返回一个初始化的空链表
func New() *Tlist {
	return new(Tlist).init()
}

// PushBack 将节点插入链表尾部
// 操作步骤:
// 1. 设置新节点的Last指向当前尾部节点
// 2. 设置新节点的Next指向哨兵节点
// 3. 修改原尾部节点的Next指向新节点
// 4. 更新哨兵节点的Last指向新节点
func (t *Tlist) PushBack(n *Node) {
	n.Last = t.root.Last // 插入的节点的next指向当前链表尾部
	n.Next = t.root      // 新插入的节点的last指向nil
	t.root.Last.Next = n // 尾部节点的next指向新插入的节点
	t.root.Last = n      // 更新尾部节点
}

// PushList 将另一个链表插入当前链表尾部
// 如果l为空或l是空链表则直接返回
// 操作完成后l链表会被清空
func (t *Tlist)PushList(l *Tlist){
	if l==nil{
		return
	}
	if l.root.Next==l.root{
		return
	}
	// 变量简化
	first,last:=l.root.Next,l.root.Last
	// 正方向插入
	t.root.Last.Next = first
	last.Next = t.root
	// 反方向插入
	first.Last=t.root.Last
	t.root.Last = last

	// 更新被改链表
	l.root.Next = l.root
	l.root.Last = l.root
}

// PushFront 将节点插入链表头部
// 操作步骤:
// 1. 设置新节点的Last指向哨兵节点
// 2. 设置新节点的Next指向当前头部节点
// 3. 修改当前头部节点的Last指向新节点
// 4. 更新哨兵节点的Next指向新节点
func (t *Tlist) PushFront(n *Node) {
	n.Last = t.root      // 新插入的节点的last指向当前链表头部
	n.Next = t.root.Next // 新插入的节点的next指向当前链表头部节点
	t.root.Next.Last = n // 更新头部节点的last
	t.root.Next = n      // 更新头部节点
}

// MoveFront2Back 将另一个链表的头部节点移动到当前链表尾部
// 操作步骤:
// 1. 将当前链表尾部节点的Next指向n的头部节点
// 2. 将n的头部节点的Last指向当前链表尾部节点
// 3. 更新当前链表的尾部节点
// 4. 更新n链表的头部节点
func (t *Tlist) MoveFront2Back(n *Tlist) {
	t.root.Last.Next = n.root.Next // 尾部节点的next指向新插入的节点
	n.root.Next.Last = t.root.Last // 插入节点的last指向尾部节点
	t.root.Last = n.root.Next      // 更新尾部节点
	n.root.Next = n.root.Next.Next // 更换n头部节点
	n.root.Next.Last = n.root      // 更新n头部节点的last
	t.root.Last.Next = n.root      // 更新插入节点的next
}

// PopFront 移除并返回链表头部节点
// 如果链表为空则返回nil
func (t *Tlist) PopFront() *Node {
	n := t.root.Next
	if n==t.root {
		tlog.Common("unusedqueue is empty", "Warning", "timewheel")
		return nil
	}
	t.root.Next = n.Next
	t.root.Next.Last = t.root
	return n
}

// PopBack 移除并返回链表尾部节点
// 如果链表为空则返回nil
func (t *Tlist) PopBack() *Node {
	n := t.root.Last
	if n==t.root {
		tlog.Common("unusedqueue is empty", "Warning", "timewheel")
		return nil
	}
	t.root.Last = n.Last
	t.root.Last.Next = t.root
	return n
}

// Delete 删除空链表
// 如果链表不为空则返回错误
// 对于build方式创建的链表会返回原队列
// 对于new方式创建的链表会直接释放
func (t *Tlist) Delete(l *Tlist) error {
	if t.root.Next != t.root {
		tlog.Common("delete fail not empty", "tlist")
		tlog.Err_in("delete fail not empty", "tlist")
		return fmt.Errorf("not empty")
	}
	l.PushBack(t.root)
	t.root=nil
	return nil
}

// Front 返回链表头部节点但不移除
// 如果链表为空则返回nil
func (t *Tlist) Front() *Node {
	if t.root.Next == t.root {
		return nil
	}
	return t.root.Next
}

// Back 返回链表尾部节点但不移除
// 如果链表为空则返回nil
func (t *Tlist) Back() *Node {
	if t.root.Last == t.root {
		return nil
	}
	return t.root.Last
}

// End 返回链表哨兵节点
// 哨兵节点既是头也是尾，用于标识链表边界
func (t *Tlist) End() *Node{
	return t.root
}