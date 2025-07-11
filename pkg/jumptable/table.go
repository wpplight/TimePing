package jumptable

import (
	"timeping/pkg/tlist"
	"timeping/pkg/upool"
)

//从跳表中取出元素
func (jt *JumpTable)Pop() uint32 {
	n:= jt.frontData()
	begin:=upool.GetDataBegin(n)
	end:=upool.GetDataEnd(n)
	if begin==end {
		jt.popData()
		unusedpool.PushDate(n)
		for i:=0;i<len(jt.indexlist);i++{
			m:=jt.popIndex(i)
			unusedpool.PushIndex(m)
		}
		return begin
	}
	upool.SetDataBegin(n,begin+1)
	for _,n:=range jt.indexlist{
		k:=n.front()
		upool.SetIndexBegin(k,begin+1)
	}

	return begin
}

//将元素归还给跳表
func (jt *JumpTable)Insert(n uint32)  {
	//获取深度和初始化栈
	l:=len(jt.indexlist)
	stack:=make([]*tlist.Node,0,l)
	
	node:=jt.indexlist[l-1].index.End()
	
	//索引层快速查找
	for i:=l-1;i>=0;i--{
		node=jt.indexlist[i].checkIndex(n,node)
		stack=append(stack,node)
		node=upool.GetIndexto(node)
	}

	//数据层移动
	w:=jt.fastCheckData(n,node)

	//后侧合并检测
	end:=upool.GetDataEnd(w)
	if end+1==n{
		upool.SetDataEnd(w,n)
		//尾部直接返回
		p:=jt.getDataEnd()
		if w.Last==p{
			return 
		}
		//需要合并
		if upool.GetDataBegin(w.Last)==n+1{
			jt.mergenode(w,stack)
		}
		return 
	}

	//添加节点
	jt.insertBackData(n,w)
	w=w.Next
	//索引层概率插入
	for i:=0;i<l;i++{
		if okinsert(){
			break
		}
		w=jt.indexlist[i].insertBackIndex(n,w,stack[l-1-i])
	}
	
	//检查更新索引层
	jt.checkLevel()
}