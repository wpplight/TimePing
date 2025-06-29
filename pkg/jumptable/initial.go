package jumptable

import (
	"timeping/pkg/tlist"
	"timeping/pkg/upool"
)


func (dl *DataList)init(begin uint32,end uint32){
	dl.Data=tlist.New();
	n:=unusedpool.PopData()
	k:=upool.NodeToDataNode(n)
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

func New(begin uint32,end uint32,unusepool *upool.Unused) (*JumpTable){

	j:=new(JumpTable)

	unusedpool=unusepool

	j.init(begin,end)

	return j
}