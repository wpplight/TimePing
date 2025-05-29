package global

import (
	"container/list"
	"timeping/tlist"

	u "timeping/utype"
)

var Trans struct{
	TaskId uintptr
	Used uintptr
}

var(
	Taskpool = list.New()
	UnuseQueue =tlist.New()
	Conf u.Cnf
	AddTaskChan = make(chan u.TaskInfo, 1)
	DeleteTaskChan = make(chan u.TaskInfo, 1)
	UpdateTaskChan = make(chan u.TaskInfo, 1)
)

