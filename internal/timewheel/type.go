package timewheel


type chanInfo struct {
	TaskId   int
	TaskPos  int
	TaskTime int
}

var (
	AddTaskChan    = make(chan chanInfo, 1)
	DeleteTaskChan = make(chan chanInfo, 1)
	UpdateTaskChan = make(chan chanInfo, 1)
)
