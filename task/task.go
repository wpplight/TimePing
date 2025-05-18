package task
import (
	"time"
)

func AddTask(TimePing []TaskNode,index int, task interface{}, time time.Time) {
	TimePing[index].task = task
	TimePing[index].time = time
	TimePing[index].isuse = true
}