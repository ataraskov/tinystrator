package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"tinystrator/task"
	"tinystrator/worker"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	host := os.Getenv("TINYSTRATOR_HOST")
	port, _ := strconv.Atoi(os.Getenv("TINYSTRATOR_PORT"))

	fmt.Println("Starting tinystrator worker")

	w := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}
	api := worker.Api{Address: host, Port: port, Worker: &w}

	go runTasks(&w)
	go w.CollectStats()
	api.Start()
}

func runTasks(w *worker.Worker) {
	for {
		if w.Queue.Len() != 0 {
			result := w.RunTask()
			if result.Error != nil {
				log.Printf("Error running task: %v\n", result.Error)
			}
		} else {
			log.Printf("No tasks to process currently.\n")
		}
		log.Println("Sleeping...")
		time.Sleep(10 * time.Second)
	}
}
