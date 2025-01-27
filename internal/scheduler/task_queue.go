package scheduler

import (
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type TaskQueue struct {
	Tasks []*Task
	mutex sync.Mutex
}

type Task struct {
	Id         string
	Expression string
	Schedule   cron.Schedule
	Command    string
	NextRun    time.Time
	IsActive   bool
}

func (q *TaskQueue) Len() int {
	return len(q.Tasks)
}

func (q *TaskQueue) Less(i, j int) bool {
	return q.Tasks[i].NextRun.Before(q.Tasks[j].NextRun)
}

func (q *TaskQueue) Swap(i, j int) {
	q.Tasks[i], q.Tasks[j] = q.Tasks[j], q.Tasks[i]
}

func (q *TaskQueue) Push(x any) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.Tasks = append(q.Tasks, x.(*Task))
}

func (q *TaskQueue) Pop() any {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	old := q.Tasks
	n := len(old)
	task := old[n-1]
	old[n-1] = nil
	q.Tasks = old[0 : n-1]
	return task
}

func (q *TaskQueue) Peek() *Task {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.Len() == 0 {
		return nil
	}

	return q.Tasks[0]
}
