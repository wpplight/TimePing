package utype

func (n *Node) Move(){
	if(n.Next==n||n==nil){
		return ;
	}
	n.Last.Next=n.Next
	n.Next.Last=n.Last
}
func (n *Node)Insert(other *Node){
	if(n== nil){
		return ;
	}
	n.Next.Last=other
	other.Next=n.Next
	other.Last=n
	n.Next=other
}