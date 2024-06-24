package worker

import (
	"fmt"
	"tinystrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("TODO: collect stats")
}

func (w *Worker) RunTask() {
	fmt.Println("TODO: start or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("TODO: start a task")
}

func (w *Worker) StopTask() {
	fmt.Println("TODO: stop a task")
}
