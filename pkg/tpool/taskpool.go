package tpool

import "timeping/pkg/tlist"

//  初始化数据池 , pagesize=2KB+2KB
func init_upool(){
	
	dupool = make([][85]DataNode, 0, 20)
	poolnode := [85]DataNode{}
	dupool = append(dupool, poolnode)
	iupool = make([][64]IndexNode, 0, 20)
	indexnode:=[64]IndexNode{}
	iupool = append(iupool, indexnode)

	create_unuesdtable()
}

// 创建一个空表
func create_unuesdtable() { 
	//初始化
	unusedpool.data=tlist.New()

	unusedpool.LenD=uint32(len(dupool[0]))

	for i:=0;i<85;i++ {
		unusedpool.data.PushBack(&dupool[0][i].Tnode)
	}

	for i:=0;i<64;i++ {
		unusedpool.index.PushBack(&iupool[0][i].Tnode)
	}
}
// 扩容Data池
func (u *unused)upData()  {
	poolnode := [85]DataNode{}
	dupool = append(dupool, poolnode)
	for i:=0;i<85;i++ {
		unusedpool.data.PushBack(&poolnode[i].Tnode)
	}
	u.LenD+=85
}

// 扩容Index池
func (u *unused)upIndex()  {
	indexnode:=[64]IndexNode{}
	iupool = append(iupool, indexnode)
	for i:=0;i<64;i++ {
		unusedpool.index.PushBack(&indexnode[i].Tnode)
	}
	u.LenI+=64
}

func (u *unused)PopIndex() *tlist.Node {
	if u.LenI==0 {
		u.upIndex()	
	}
	k:=u.index.PopFront()
	u.LenI--
	return k
}

func (u *unused)PopData() *tlist.Node {
	if u.LenD==0 {
		u.upData()
	}
	k:=u.data.PopFront()
	u.LenD--
	return k
}

func (u *unused)pushDate(node *tlist.Node){
	u.data.PushBack(node)
	u.LenD++
}
func (u *unused)pushIndex(node *tlist.Node){
	u.index.PushBack(node)
	u.LenI++
}