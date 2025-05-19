package tlist
import(
	"timeping/utype"
)
type Tlist struct{
	len int
	head *utype.Node
	tail *utype.Node
}
func (t *Tlist) init() *Tlist {
	
	t.head = nil
	t.tail = nil
	t.len = 0	
	return t
}

//创建侵入式链表,会返回一个该链表的指针
func New() *Tlist{
	return new(Tlist).init()
}
//返回该链表长度
func (t *Tlist) Length() int {
	return t.len
}
//输入节点指针，就会被推入尾部
func (t *Tlist) PushBack(n *utype.Node) { 
	t.tail  = n
	n.Last  = t.head
	t.len++
}

func (t *Tlist) PushFront(n *utype.Node) { 
	t.head  = n
	n.Next  = t.tail
	t.len++
}
func (t *Tlist) MoveFront2Back(n Tlist) { 
	
}


func (t *Tlist) PopFront() *utype.Node { 
	if t.len == 0 {
		return nil
	}
	n := t.head
	t.head = n.Next
	t.len--
	return n
}
func (t *Tlist) PopBack() *utype.Node { 
	if t.len == 0 {
		return nil
	}
	n := t.tail
	t.tail = n.Last
	t.len--
	return n
}