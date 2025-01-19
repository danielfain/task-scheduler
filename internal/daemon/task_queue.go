package daemon

import "task-scheduler/internal/common/types"

type TaskQueue []*types.Task

func (q TaskQueue) Len() int {
	return len(q)
}

func (q TaskQueue) Less(i, j int) bool {
	return q[i].ScheduledTime.UnixMilli() < q[j].ScheduledTime.UnixMilli()
}

func (q TaskQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].QueueIndex = i
	q[j].QueueIndex = j
}

func (q *TaskQueue) Push(x any) {
	n := len(*q)
	task := x.(*types.Task)
	task.QueueIndex = n
	*q = append(*q, task)
}

func (q *TaskQueue) Pop() any {
	old := *q
	n := len(old)
	task := old[n-1]
	old[n-1] = nil
	task.QueueIndex = -1
	*q = old[0 : n-1]
	return task
}
