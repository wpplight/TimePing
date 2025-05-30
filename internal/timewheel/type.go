package timewheel

import "timeping/pkg/tlist"
var(
	AddTaskChan = make(chan tlist.Node, 1)
	DeleteTaskChan = make(chan tlist.Node, 1)
	UpdateTaskChan = make(chan tlist.Node, 1)
)

type Twhell struct{
	T *tlist.Tlist
	N tlist.Node
}