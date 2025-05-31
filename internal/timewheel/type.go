package timewheel

import "time"

type chanInfo struct {
	TaskId   uint16
	TaskPos  int
	TaskTime int
}

var (
	AddTaskChan    = make(chan chanInfo, 1)
	DeleteTaskChan = make(chan chanInfo, 1)
	UpdateTaskChan = make(chan chanInfo, 1)
)
