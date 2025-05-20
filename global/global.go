package global

import (
	"container/list"
	"timeping/tlist"

	u "timeping/utype"
)

var Trans struct{
	TCB uintptr
	
}

var(
	Taskpool = list.New()
	UnuseQueue =tlist.New()
	Conf u.Cnf
	TaskChan = make(chan u.TaskInfo, 1)
)

