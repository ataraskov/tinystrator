package manager

import (
	"fmt"
	"tinystrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("TODO: select an appropriate worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("TODO: update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("TODO: send work to workers")
}
