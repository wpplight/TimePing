package task


func creatUnusedTask(taskPoolNode []TaskNode) {

	for i:=0;i<len(taskPoolNode);i++{

		taskPoolNode[i].Used = false

		taskPoolNode[i].TaskId  = uint16(i)

		UnuseQueue.PushBack(&taskPoolNode[i].Tnode)
	}

}