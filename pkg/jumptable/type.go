package jumptable

import (
	"timeping/pkg/tlist"
	"timeping/pkg/upool"
)

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

// 数据队列结构体 16B
type DataList struct {
	length uint64 //8byte
	Data *tlist.Tlist //8byte
}




var (
	unusedpool *upool.Unused// 未使用的数据池
	deepln int
)