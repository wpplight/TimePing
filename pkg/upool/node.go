package upool

import (
	"timeping/pkg/tlist"
	"unsafe"
)

// Reset 重置 DataNode 的开始和结束值
func (dn *DataNode) Reset(begin uint32, end uint32){
	dn.Begin = begin
	dn.End = end
}

// NodeToDataNode 将 tlist.Node 转换为 DataNode 类型
func NodeToDataNode(n *tlist.Node) *DataNode{
	return (*DataNode)(unsafe.Pointer(uintptr(unsafe.Pointer(n))+dn))
}

// NodeToIndex 将 tlist.Node 转换为 IndexNode 类型
func NodeToIndex(n *tlist.Node) *IndexNode{
	return (*IndexNode)(unsafe.Pointer(uintptr(unsafe.Pointer(n))+in))
}

// setIndexBegin 设置 IndexNode 的索引值
func SetIndexBegin(n *tlist.Node,begin uint32){
	k:=NodeToIndex(n)
	k.Begin = begin
}

// setDataBegin 设置 DataNode 的开始值
func SetDataBegin(n *tlist.Node,begin uint32){
	k:=NodeToDataNode(n)
	k.Begin = begin
}

// setDataEnd 设置 DataNode 的结束值
func SetDataEnd(n *tlist.Node,end uint32){
	k:=NodeToDataNode(n)
	k.End = end
}

// setIndexto 设置 IndexNode 的指向节点
func SetIndexto(n *tlist.Node,to *tlist.Node){
	k:=NodeToIndex(n)
	k.To = to
}

// getIndexEnd 获取 IndexNode 的结束值
func GetIndexBegin(n *tlist.Node) uint32{
	k:=NodeToIndex(n)
	return k.Begin
}

// getDataBegin 获取 DataNode 的开始值
func GetDataBegin(n *tlist.Node) uint32{
	k:=NodeToDataNode(n)
	return k.Begin
}

// getDataEnd 获取 DataNode 的结束值
func GetDataEnd(n *tlist.Node) uint32{
	k:=NodeToDataNode(n)
	return k.End
}

// getIndexto 获取 IndexNode 的指向节点
func GetIndexto(n *tlist.Node) *tlist.Node{
	k:=NodeToIndex(n)
	return k.To
}