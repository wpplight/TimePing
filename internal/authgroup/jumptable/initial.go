package jumptable

import (
	"timeping/pkg/tlist"
	"unsafe"
)


func (dl *DataList)init(begin uint32,end uint32){
	dl.Data=tlist.New();
	n:=unusedpool.PopData()
	k:=NodeToDataNode(n)
	k.Reset(begin,end)
	dl.Data.PushBack(n)
	dl.length=1
}

func (j *JumpTable)init(begin uint32,end uint32){ 

	j.indexlist  = make([]Index,0,20)
	j.indexlist=  append(j.indexlist,Index{1,tlist.New()})

	j.data = new(DataList)
	j.data.init(begin,end)
	j.addlevel()
	deepln=2
	
}

func New(begin uint32,end uint32) (*JumpTable){

	j:=new(JumpTable)
	init_upool()
	
	dn=unsafe.Offsetof(DataNode{}.Tnode)
	in=unsafe.Offsetof(IndexNode{}.Tnode)

	j.init(begin,end)

	return j
}