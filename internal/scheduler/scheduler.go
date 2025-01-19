package scheduler

import (
	"container/heap"
	"time"

	"github.com/robfig/cron/v3"
)

var queue = make(TaskQueue, 0)

func ScheduleTask(cronExpression string, command string) {
	task := &Task{
		State:          SCHEDULED,
		CronExpression: cronExpression,
		Command:        command,
		ScheduledTime:  nextScheduledTime(cronExpression),
	}

	heap.Push(&queue, task)
}

func GetTasks() TaskQueue {
	return queue
}

func nextScheduledTime(cronExpression string) time.Time {
	schedule, err := cron.ParseStandard(cronExpression)
	if err != nil {
		panic(err.Error())
	}

	return schedule.Next(time.Now())
}
