package upool

import "timeping/pkg/tlist"

type Unused struct{
	LenD uint32
	LenI uint32

	data *tlist.Tlist
	index *tlist.Tlist
}

//  数据节点结构体 24B
type DataNode struct {
	Begin uint32 //4byte	 
	End uint32 //4byte
	Tnode tlist.Node // 8+8byte
}

// 索引节点结构体 32B
type IndexNode struct {
	Begin uint32 //4byte
	To *tlist.Node	// 8byte
	Tnode tlist.Node // 8+8byte
}

var (
	dn uintptr //tnode到dataNode的偏移
	in uintptr //tnode到indexNode的偏移
	
	dupool [][85]DataNode // datanode的数据池
	iupool [][64]IndexNode // indexnode的数据池

	Unusedpool Unused// 未使用的数据池
)