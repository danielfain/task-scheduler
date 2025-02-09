package daemon

import (
	"context"
	"database/sql"
	"log"
	"net"
	"task-scheduler/proto"
	"time"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskSchedulerServer struct {
	proto.UnimplementedTaskSchedulerServer
	db    *sql.DB
	queue *TaskQueue
}

func (s *TaskSchedulerServer) AddTask(ctx context.Context, req *proto.AddTaskRequest) (*proto.TaskId, error) {
	schedule, err := cron.ParseStandard(req.Expression)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid cron expression: %v", err)
	}

	taskId := uuid.New().String()
	nextRun := schedule.Next(time.Now())
	_, err = s.db.Exec(`
		insert into tasks (id, expression, command, next_run, is_active)
		values (?, ?, ?, ?, ?)
		`, taskId, req.Expression, req.Command, nextRun, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error inserting task: %v", err)
	}

	s.queue.Push(&Task{
		Id:         taskId,
		Expression: req.Expression,
		Schedule:   schedule,
		Command:    req.Command,
		NextRun:    nextRun,
		IsActive:   true,
	})

	return &proto.TaskId{Id: taskId}, nil
}

func (s *TaskSchedulerServer) ListTasks(empty *proto.Empty, stream proto.TaskScheduler_ListTasksServer) error {
	for _, task := range s.queue.tasks {
		stream.Send(&proto.Task{
			Id:         task.Id,
			Expression: task.Expression,
			Command:    task.Command,
			NextRun:    task.NextRun.String(),
			IsActive:   task.IsActive,
		})
	}

	return nil
}

func StartServer(db *sql.DB) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTaskSchedulerServer(grpcServer, &TaskSchedulerServer{db: db})
	log.Printf("gRPC server listening on %s", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startScheduler(q *TaskQueue) {
	for {
		if job := q.Peek(); job != nil {
			now := time.Now()
			waitTime := job.NextRun.Sub(now)

			if waitTime <= 0 {
				// Execute the job
				go executeJob(job)
				// Reschedule the job for next run
				nextRun := cronParser.Parse(job.Schedule).Next(now)
				q.RescheduleJob(job.ID, nextRun)
			} else {
				// Wait until the next job is ready
				time.Sleep(waitTime)
			}
		} else {
			// No jobs—sleep briefly to avoid CPU spin
			time.Sleep(1 * time.Second)
		}
	}
}
