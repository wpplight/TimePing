package upool

import (
	"timeping/pkg/tlist"
	"unsafe"
)

//  初始化数据池 , pagesize=2KB+2KB
func Init_upool(){
	
	dupool = make([][85]DataNode, 0, 20)
	poolnode := [85]DataNode{}
	dupool = append(dupool, poolnode)
	iupool = make([][64]IndexNode, 0, 20)
	indexnode:=[64]IndexNode{}
	iupool = append(iupool, indexnode)
	dn=unsafe.Offsetof(DataNode{}.Tnode)
	in=unsafe.Offsetof(IndexNode{}.Tnode)

	create_unuesdtable()
}

// 创建一个空表
func create_unuesdtable() { 
	//初始化
	Unusedpool.data=tlist.New()
	Unusedpool.index=tlist.New()

	Unusedpool.LenD=uint32(len(dupool[0]))
	Unusedpool.LenI=uint32(len(iupool[0]))

	for i:=0;i<85;i++ {
		Unusedpool.data.PushBack(&dupool[0][i].Tnode)
	}

	for i:=0;i<64;i++ {
		Unusedpool.index.PushBack(&iupool[0][i].Tnode)
	}
}
// 扩容Data池
func (u *Unused)upData()  {
	poolnode := [85]DataNode{}
	dupool = append(dupool, poolnode)
	for i:=0;i<85;i++ {
		Unusedpool.data.PushBack(&poolnode[i].Tnode)
	}
	u.LenD+=85
}

// 扩容Index池
func (u *Unused)upIndex()  {
	indexnode:=[64]IndexNode{}
	iupool = append(iupool, indexnode)
	for i:=0;i<64;i++ {
		Unusedpool.index.PushBack(&indexnode[i].Tnode)
	}
	u.LenI+=64
}

func (u *Unused)PopIndex() *tlist.Node {
	if u.LenI==0 {
		u.upIndex()	
	}
	k:=u.index.PopFront()
	u.LenI--
	return k
}

func (u *Unused)PopData() *tlist.Node {
	if u.LenD==0 {
		u.upData()
	}
	k:=u.data.PopFront()
	u.LenD--
	return k
}

func (u *Unused)PushDate(node *tlist.Node){
	u.data.PushBack(node)
	u.LenD++
}
func (u *Unused)PushIndex(node *tlist.Node){
	u.index.PushBack(node)
	u.LenI++
}