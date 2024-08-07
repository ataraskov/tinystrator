package main

import (
	"fmt"
	"os"
	"strconv"
	"tinystrator/manager"
	"tinystrator/task"
	"tinystrator/worker"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	mhost := os.Getenv("TINYSTRATOR_MANAGER_HOST")
	mport, _ := strconv.Atoi(os.Getenv("TINYSTRATOR_MANAGER_PORT"))

	whost := os.Getenv("TINYSTRATOR_WORKER_HOST")
	wport, _ := strconv.Atoi(os.Getenv("TINYSTRATOR_WORKER_PORT"))

	fmt.Println("Starting tinystrator worker")
	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}
	wapi := worker.Api{Address: whost, Port: wport, Worker: &w}

	go w.RunTasks()
	go w.UpdateTasks()
	go w.CollectStats()
	go wapi.Start()

	fmt.Println("Starting tinystrator manager")
	workers := []string{fmt.Sprintf("%s:%d", whost, wport)}
	m := manager.New(workers)
	mapi := manager.Api{Address: mhost, Port: mport, Manager: m}

	go m.UpdateTasks()
	go m.ProcessTasks()
	go m.DoHealthChecks()
	mapi.Start()
}
