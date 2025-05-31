package tlist

//tlist中的基本单元纯粹的链表节点，存储了上一个节点和下一个节点的地址
type Node struct{
	Next *Node
	Last *Node
}