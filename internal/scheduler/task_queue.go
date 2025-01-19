package scheduler

type TaskQueue []*Task

func (q TaskQueue) Len() int {
	return len(q)
}

func (q TaskQueue) Less(i, j int) bool {
	return q[i].ScheduledTime.UnixMilli() < q[j].ScheduledTime.UnixMilli()
}

func (q TaskQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *TaskQueue) Push(x any) {
	n := len(*q)
	task := x.(*Task)
	task.index = n
	*q = append(*q, task)
}

func (q *TaskQueue) Pop() any {
	old := *q
	n := len(old)
	task := old[n-1]
	old[n-1] = nil
	task.index = -1
	*q = old[0 : n-1]
	return task
}
