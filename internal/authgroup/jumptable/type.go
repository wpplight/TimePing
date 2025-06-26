package jumptable

import "timeping/pkg/tlist"

// 跳表对象 
type JumpTable struct {
	data *DataList //8byte
	indexlist []Index  //16byte
}
//  索引队列结构体 16B
type Index struct{
	length uint64	//8byte
	index *tlist.Tlist // 8byte
}
// 索引节点结构体 32B
type IndexNode struct {
	Begin uint32 //4byte
	To *tlist.Node	// 8byte
	Tnode tlist.Node // 8+8byte
}
// 数据队列结构体 16B
type DataList struct {
	length uint64 //8byte
	Data *tlist.Tlist //8byte
}
//  数据节点结构体 24B
type DataNode struct {
	Begin uint32 //4byte	 
	End uint32 //4byte
	Tnode tlist.Node // 8+8byte
}


type unused struct{
	LenD uint32
	LenI uint32

	data *tlist.Tlist
	index *tlist.Tlist
}

var (
	dn uintptr //tnode到dataNode的偏移
	in uintptr //tnode到indexNode的偏移

	dupool [][85]DataNode // datanode的数据池
	iupool [][64]IndexNode // indexnode的数据池

	unusedpool unused// 未使用的数据池

	deepln int
)