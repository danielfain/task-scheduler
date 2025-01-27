package main

import (
	"container/heap"
	"log"
	"task-scheduler/internal/scheduler"
	"task-scheduler/internal/server"
	"task-scheduler/internal/storage"
	"time"
)

func main() {
	db, err := storage.InitDB()
	if err != nil {
		log.Fatalf("could not create db: %v", err)
	}
	defer db.Close()

	q := &scheduler.TaskQueue{}
	heap.Init(q)
	if err := storage.LoadTasks(db, q); err != nil {
		log.Fatalf("Failed to load tasks: %v", err)
	}

	go server.StartGRPCServer(db, q)

	for {
		if job := q.Peek(); job != nil {
			now := time.Now()
			waitTime := job.NextRun.Sub(now)

			if waitTime <= 0 {
				// TODO: Execute and recalculate next run time
			} else {
				time.Sleep(waitTime)
			}
		} else {
			// No jobs—sleep briefly to avoid CPU spin
			time.Sleep(1 * time.Second)
		}
	}
}
