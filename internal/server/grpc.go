package server

import (
	"context"
	"database/sql"
	"log"
	"net"
	"task-scheduler/internal/scheduler"
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
	queue *scheduler.TaskQueue
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

	s.queue.Push(&scheduler.Task{
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
	for _, task := range s.queue.Tasks {
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

func StartGRPCServer(db *sql.DB, queue *scheduler.TaskQueue) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTaskSchedulerServer(grpcServer, &TaskSchedulerServer{db: db, queue: queue})
	log.Printf("gRPC server listening on %s", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
