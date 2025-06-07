package jumptable

type JumpTable struct {
    
}
type DataNode struct {
	Begin uint16
	End uint16
	next *DataNode
}