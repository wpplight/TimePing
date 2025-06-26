package tlist

// Move 将当前节点从双向链表中移除
// 方法逻辑：
// 1. 如果节点是空指针或自身循环(Next指向自己)，直接返回
// 2. 否则调整相邻节点的指针，跳过当前节点
func (n *Node) Move(){
	if(n.Next==n||n==nil){
		return ;
	}
	n.Last.Next=n.Next
	n.Next.Last=n.Last
}

// InsertBack 将other节点插入到当前节点之后
// 方法逻辑：
// 1. 如果当前节点是空指针，直接返回
// 2. 将other节点插入到当前节点和原后继节点之间
// 3. 调整三个节点之间的指针关系
func (n *Node)InsertBack(other *Node){
	if(n== nil){
		return ;
	}
	n.Next.Last=other
	other.Next=n.Next
	other.Last=n
	n.Next=other
}

// InsertFront 将other节点插入到当前节点之前
// 方法逻辑：
// 1. 如果当前节点是空指针，直接返回
// 2. 将other节点插入到当前节点和原前驱节点之间
// 3. 调整三个节点之间的指针关系
func (n * Node)InsertFront(other *Node){
	if(n==nil){
		return ;
	}
	n.Last.Next,other.Last=other,n.Last
	other.Next,n.Last=n,other
}