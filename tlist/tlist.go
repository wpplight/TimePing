package tlist

import (
	"fmt"
	"timeping/tlog"
	"timeping/utype"
)

type Tlist struct {
	origin *Tlist      //归属队列
	root   *utype.Node //哨兵节点
}

func (t *Tlist) init() *Tlist {
	t.root = new(utype.Node)
	t.root.Last = t.root //指向队尾
	t.root.Next = t.root //指向队首
	t.origin = nil
	return t
}

// 使用已有内存块创建一个链表，并返回一个该链表的指针
func Build(n *Tlist) (*Tlist, error) {
	if n == nil {
		tlog.Common("Tlist build error nil", "tlist")
		return nil, fmt.Errorf("Tlist init error")
	}
	k := new(Tlist)
	if n.root.Next == n.root {
		tlog.Common("Tlist build error few element", "tlist")
		return nil, fmt.Errorf("Tlist init error")
	}
	k.root = n.PopFront()
	k.origin = n
	return k, nil
}

// 创建侵入式链表,会返回一个该链表的指针
func New() *Tlist {
	return new(Tlist).init()
}

// 输入节点指针，就会被推入尾部
func (t *Tlist) PushBack(n *utype.Node) {
	n.Last = t.root.Last //插入的节点的next指向当前链表尾部
	n.Next = t.root      //新插入的节点的last指向nil
	t.root.Last.Next = n //尾部节点的next指向新插入的节点
	t.root.Last = n      //更新尾部节点
}

func (t *Tlist) PushFront(n *utype.Node) {
	n.Last = t.root      //新插入的节点的last指向当前链表头部
	n.Next = t.root.Next //新插入的节点的next指向当前链表头部节点
	t.root.Next.Last = n //更新头部节点的last
	t.root.Next = n      //更新头部节点
}
func (t *Tlist) MoveFront2Back(n *Tlist) {
	t.root.Last.Next = n.root.Next //尾部节点的next指向新插入的节点
	n.root.Next.Last = t.root.Last //插入节点的last指向尾部节点
	t.root.Last = n.root.Next      //更新尾部节点
	n.root.Next = n.root.Next.Next //更换n头部节点
	n.root.Next.Last = n.root      //更新n头部节点的last
	t.root.Last.Next = n.root      //更新插入节点的next
}

// 返回链表头部节点,并删除
func (t *Tlist) PopFront() *utype.Node {
	n := t.root.Next
	t.root.Next = n.Next
	t.root.Next.Last = t.root
	return n
}

// 返回链表尾部节点,并删除
func (t *Tlist) PopBack() *utype.Node {
	n := t.root.Last
	t.root.Last = n.Last
	t.root.Last.Next = t.root
	return n
}
