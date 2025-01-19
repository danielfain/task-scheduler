package daemon

import (
	"container/heap"
	"task-scheduler/internal/common/types"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	queue TaskQueue
}

func NewScheduler() Scheduler {
	return Scheduler{
		queue: make(TaskQueue, 0),
	}
}

func (s Scheduler) ScheduleTask(cronExpression string, command string) {
	task := &types.Task{
		State:          types.SCHEDULED,
		CronExpression: cronExpression,
		Command:        command,
		ScheduledTime:  nextScheduledTime(cronExpression),
	}

	heap.Push(&s.queue, task)
}

func nextScheduledTime(cronExpression string) time.Time {
	schedule, err := cron.ParseStandard(cronExpression)
	if err != nil {
		panic(err.Error())
	}

	return schedule.Next(time.Now())
}
