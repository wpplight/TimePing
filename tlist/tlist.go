package tlist

import (
	"timeping/tlog"
	"timeping/utype"
	"fmt"
)
//头节点
type Tlist struct{
	len int
	root *utype.Node //哨兵节点
}
func (t *Tlist) init() *Tlist {
	t.root = new(utype.Node)
	t.root.Last =t.root //指向队尾
	t.root.Next =t.root //指向队首
	t.len = 0	
	return t
}

// 使用已有内存块创建一个链表，并返回一个该链表的指针
func Build(n *utype.Node) (*Tlist,error) {
	if n == nil {
		tlog.Common("Tlist init error","tlist")
		return nil,fmt.Errorf("Tlist init error")
	}
	k:=new(Tlist)
	k.root= n
	return k,nil
}

//创建侵入式链表,会返回一个该链表的指针
func New() *Tlist{
	return new(Tlist).init()
}        

//返回该链表长度
func (t *Tlist) Length() int {                     
	return t.len
}
//输入节点指针，就会被推入尾部
func (t *Tlist) PushBack(n *utype.Node) { 
	n.Last = t.root.Last //插入的节点的next指向当前链表尾部
	n.Next=t.root	//新插入的节点的last指向nil
	t.root.Last.Next = n	//尾部节点的next指向新插入的节点
	t.root.Last = n		//更新尾部节点
	t.len++
}   

//输入节点指针，就会被推入头部
func (t *Tlist) PushFront(n *utype.Node) { 
	n.Last =t.root	//新插入的节点的last指向当前链表头部
	n.Next = t.root.Next	//新插入的节点的next指向当前链表头部节点
	t.root.Next.Last = n	//更新头部节点的last
	t.root.Next = n	//更新头部节点
	t.len++
}
func (t *Tlist) MoveFront2Back(n *Tlist) { 
	if(n.len==0){
		return
	}
	n.len--
	t.root.Last.Next = n.root.Next	//尾部节点的next指向新插入的节点
	n.root.Next.Last = t.root.Last	//插入节点的last指向尾部节点
	t.root.Last = n.root.Next //更新尾部节点
	n.root.Next =n.root.Next.Next	//更换n头部节点
	n.root.Next.Last =n.root	//更新n头部节点的last
	t.root.Last.Next =n.root	//更新插入节点的next
	t.len++
}

//返回链表头部节点,并删除
func (t *Tlist) PopFront() *utype.Node { 
	if t.len == 0 {
		return nil
	}
	n := t.root.Next
	t.root.Next = n.Next
	t.root.Next.Last =t.root
	t.len--
	return n
}
//返回链表尾部节点,并删除
func (t *Tlist) PopBack() *utype.Node { 
	if t.len == 0 {
		return nil
	}
	n := t.root.Last
	t.root.Last = n.Last
	t.root.Last.Next=t.root
	t.len--
	return n
}

func (t *Tlist) GetHeadNode() *utype.Node{
	return t.root.Next
}

func (t *Tlist) GetTailNode() *utype.Node{
	return t.root.Last
}