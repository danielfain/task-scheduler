package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"task-scheduler/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewTaskSchedulerClient(conn)

	switch os.Args[1] {
	case "add":
		taskId, err := client.AddTask(context.Background(), &proto.AddTaskRequest{
			Expression: "* 5 * * *",
			Command:    "ls",
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Created task with id: %s", taskId.Id)
	case "list":
		stream, err := client.ListTasks(context.Background(), &proto.Empty{})
		if err != nil {
			log.Fatal(err)
		}

		for {
			task, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			log.Println(task)
		}

	default:
		fmt.Println("Usage: cli [add|list|remove] ...")
	}
}
