package task
import (
	"time"
	"timeping/config"
)
type TaskNode struct {
	taskId int
	task interface{}
	isuse bool
	node *TaskNode
	next *TaskNode
	last *TaskNode
	time time.Time
}
func InitialTaskPool() []TaskNode {
	TaskPool := make([]TaskNode, config.TaskPoolSize)
	return TaskPool
}