package scheduler

import "time"

type Task struct {
	State          TaskState
	CronExpression string
	Command        string
	ScheduledTime  time.Time

	index int
}

type TaskState string

const (
	SCHEDULED  TaskState = "Scheduled"
	RUNNING    TaskState = "Running"
	TERMINATED TaskState = "Terminated"
)
