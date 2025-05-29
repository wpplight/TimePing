package tlist

//将节点从链表中删除
func (n *Node) Move(){
	if(n.Next==n||n==nil){
		return ;
	}
	n.Last.Next=n.Next
	n.Next.Last=n.Last
}
//将节点插在某个几点后面
func (n *Node)Insert(other *Node){
	if(n== nil){
		return ;
	}
	n.Next.Last=other
	other.Next=n.Next
	other.Last=n
	n.Next=other
}