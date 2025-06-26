package jumptable

import (
	"timeping/pkg/tlist"
)


func (jt *JumpTable)addlevel(){

	jt.indexlist=append(jt.indexlist, Index{0,tlist.New()})
	l:=len(jt.indexlist)-1
	if l==0{
		n:=unusedpool.PopIndex()
		d:=jt.frontData()
		setIndexBegin(n,getDataBegin(d))
		setIndexto(n,d)
		jt.indexlist[0].pushBack(n)
		return 
	}

	node:=jt.indexlist[l-1].front()
	n:=unusedpool.PopIndex()
	setIndexBegin(n,getIndexBegin(node))
	setIndexto(n,node)
	jt.indexlist[l].pushBack(n)

	for node=node.Next;node!=jt.indexlist[l-1].index.End();node=node.Next{
		if okinsert(){
			n:=unusedpool.PopIndex()
			setIndexBegin(n,getIndexBegin(node))
			setIndexto(n,node)
			jt.indexlist[l].pushBack(n)
		}
	}
} 

func (jt *JumpTable) popData() *tlist.Node {
	if jt.data.length==0{
		return nil
	}
	k:=jt.data.Data.PopFront()
	jt.data.length--
	return k
}

func (jt *JumpTable) frontData() *tlist.Node { 
	if jt.data.length==0{
		return nil
	}
	return jt.data.Data.Front()
}

func (jt *JumpTable) frontIndex(index int) *tlist.Node { 
	return jt.indexlist[index].front()
}

func (jt *JumpTable)popIndex(index int) *tlist.Node { 
	if jt.indexlist[index].length==0{
		return nil
	}
	k:=jt.indexlist[index].index.PopFront()
	jt.indexlist[index].length--
	return k
}


func (jt *JumpTable)insertFrontData(num uint32,where *tlist.Node){
	jt.data.length++
	n:=unusedpool.PopData()
	setDataBegin(n,num)
	setDataEnd(n,num)
	where.InsertFront(n)
} 
func (jt *JumpTable)insertBackData(num uint32,where *tlist.Node){
	jt.data.length++
	n:=unusedpool.PopData()
	setDataBegin(n,num)
	setDataEnd(n,num)
	where.InsertBack(n)
}

func(jt *JumpTable)fastCheckData(num uint32,where *tlist.Node) *tlist.Node { 
	end:=jt.data.Data.End()
	
	for  n:=where;n!=end;n=n.Next {
		if getDataBegin(n)>num {
			return n
		}
	}

	return jt.data.Data.Back()
}

func (jt *JumpTable)deleteData(where *tlist.Node){ 
	jt.data.length--
	where.Move()
	unusedpool.pushDate(where)
}

func (jt *JumpTable)getDataEnd() *tlist.Node{
	return jt.data.Data.End()
}

func (jt *JumpTable)mergenode(node *tlist.Node,stack []*tlist.Node){
	//删除data合并部分
	next:=node.Next
	setDataEnd(node,getDataEnd(next))
	begin:=getDataBegin(node)
	jt.deleteData(next)

	//index层便利删除废弃索引
	l:=len(stack)
	for i:=0;i<l;i++{
		n:=stack[l-1-i].Next
		if n==jt.indexlist[i].index.End(){
			break
		}	
		if getIndexBegin(n)!=begin{
			break
		}
		jt.indexlist[i].move(n)
	}

}