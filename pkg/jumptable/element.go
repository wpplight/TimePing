package jumptable

import (
	"timeping/pkg/tlist"
	"timeping/pkg/upool"
)


func (jt *JumpTable)addlevel(){

	jt.indexlist=append(jt.indexlist, Index{0,tlist.New()})
	l:=len(jt.indexlist)-1
	if l==0{
		n:=unusedpool.PopIndex()
		d:=jt.frontData()
		upool.SetIndexBegin(n,upool.GetDataBegin(d))
		upool.SetIndexto(n,d)
		jt.indexlist[0].pushBack(n)
		return 
	}

	node:=jt.indexlist[l-1].front()
	n:=unusedpool.PopIndex()
	upool.SetIndexBegin(n,upool.GetIndexBegin(node))
	upool.SetIndexto(n,node)
	jt.indexlist[l].pushBack(n)

	for node=node.Next;node!=jt.indexlist[l-1].index.End();node=node.Next{
		if okinsert(){
			n:=unusedpool.PopIndex()
			upool.SetIndexBegin(n,upool.GetIndexBegin(node))
			upool.SetIndexto(n,node)
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
	upool.SetDataBegin(n,num)
	upool.SetDataEnd(n,num)
	where.InsertFront(n)
} 
func (jt *JumpTable)insertBackData(num uint32,where *tlist.Node){
	jt.data.length++
	n:=unusedpool.PopData()
	upool.SetDataBegin(n,num)
	upool.SetDataEnd(n,num)
	where.InsertBack(n)
}

func(jt *JumpTable)fastCheckData(num uint32,where *tlist.Node) *tlist.Node { 
	end:=jt.data.Data.End()
	
	for  n:=where;n!=end;n=n.Next {
		if upool.GetDataBegin(n)>num {
			return n
		}
	}

	return jt.data.Data.Back()
}

func (jt *JumpTable)deleteData(where *tlist.Node){ 
	jt.data.length--
	where.Move()
	unusedpool.PushDate(where)
}

func (jt *JumpTable)getDataEnd() *tlist.Node{
	return jt.data.Data.End()
}

func (jt *JumpTable)mergenode(node *tlist.Node,stack []*tlist.Node){
	//删除data合并部分
	next:=node.Next
	upool.SetDataEnd(node,upool.GetDataEnd(next))
	begin:=upool.GetDataBegin(node)
	jt.deleteData(next)

	//index层便利删除废弃索引
	l:=len(stack)
	for i:=0;i<l;i++{
		n:=stack[l-1-i].Next
		if n==jt.indexlist[i].index.End(){
			break
		}	
		if upool.GetIndexBegin(n)!=begin{
			break
		}
		jt.indexlist[i].move(n)
	}

}