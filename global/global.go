package global

import (
	"container/list"
	"timeping/tlist"
	u "timeping/utype"
)

var UnuseQueue =tlist.New()
var Taskpool = list.New()
var(
	Conf u.Cnf
	TaskChan = make(chan u.TaskInfo, 1)
)

