package jumptable

import (
	"math/rand"
	"timeping/pkg/tlist"
)

// okinsert 根据50%概率决定是否执行插入操作
//
// 返回: true表示执行插入，false表示不执行插入
func okinsert() bool {
	n := rand.Intn(2)
	return n == 1
}

// pushBack 将节点添加到索引列表末尾
//
// n: 要添加的节点指针
func (il *Index) pushBack(n *tlist.Node) {
	il.index.PushBack(n)
	il.length++
}

// front 获取索引列表的第一个节点
//
// 返回: 列表首节点指针，如果列表为空则返回nil
func (il *Index) front() *tlist.Node {
	return il.index.Front()
}

// back 获取索引列表的最后一个节点
//
// 返回: 列表尾节点指针，如果列表为空则返回nil
func (il *Index) back() *tlist.Node {
	return il.index.Back()
}

// move 从索引列表中移除指定节点
//
// n: 要移除的节点指针
func (il *Index) move(n *tlist.Node) {
	n.Move()
	il.length--
}

// IsEmpty 检查索引列表是否为空
//
// 返回: true表示列表为空，false表示列表不为空
func (il *Index) IsEmpty() bool {
	return il.length == 0
}

// popFront 移除并返回索引列表的第一个节点
//
// 返回: 被移除的首节点指针
func (il *Index) popFront() *tlist.Node {
	il.length--
	return il.index.PopFront()
}

// checkIndex 检查并返回适合插入指定索引值的节点位置
//
// i: 目标索引值
// node: 起始检查节点
//
// 返回: 适合插入位置的节点指针
func (l *Index) checkIndex(i uint32, node *tlist.Node) *tlist.Node {

	list := l.index
	for n := node; n.Last != list.End(); n = n.Last {
		t := getIndexBegin(n.Last)
		if t>i {
			return n
		}
	}
	return list.Back()
}

// insertFrontIndex 从内存池获取索引节点并插入到指定节点前
//
// num: 索引值
// to: 目标节点指针
// where: 插入位置节点指针
func (i *Index) insertFrontIndex(num uint32, to *tlist.Node, where *tlist.Node) {
	n := unusedpool.PopIndex()
	i.length++
	setIndexBegin(n, num)
	setIndexto(n, to)
	where.InsertFront(n)
}

// insertBackIndex 从内存池获取索引节点并插入到指定节点后
//
// num: 索引值
// to: 目标节点指针
// where: 插入位置节点指针
func (i *Index) insertBackIndex(num uint32, to *tlist.Node, where *tlist.Node) *tlist.Node {
	n := unusedpool.PopIndex()
	i.length++
	setIndexBegin(n, num)
	setIndexto(n, to)
	where.InsertBack(n)
	return n
}

// fastInsert 快速插入节点后
//
// num: 要插入的索引值
// where: 插入位置参考节点指针
func (i *Index) fastInsert(num uint32, where *tlist.Node) {
	other := getIndexBegin(where)
	if other > num {
		i.insertFrontIndex(num, where.Last, where)
	} else {
		i.insertBackIndex(num, where.Last, where)
	}
}

func (i *Index)end() *tlist.Node {
	return i.index.End()
}


func (jt *JumpTable)checkLevel() {
	if jt.data.length < uint64(deepln){
		return 
	}
	
	l:=jt.indexlist[len(jt.indexlist)-1]
	p:=Index{0,tlist.New()}
	node:=unusedpool.PopIndex()
	setIndexBegin(node,getIndexBegin(l.front()))
	setIndexto(node,l.front())
	p.pushBack(node)
	for n:=l.front().Next;n!=l.end();n=n.Next{
		if okinsert(){
			continue
		}
		temp:=unusedpool.PopIndex()
		no:=NodeToIndex(temp) 
		no.Begin=getIndexBegin(n)
		no.To=n
		p.pushBack(temp)
	}
	deepln*=2
}

