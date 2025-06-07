package jumptable

func (dn *DataNode) Reset(begin uint16, end uint16){
	dn.Begin = begin
	dn.End = end
}