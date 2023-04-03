package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter5/proto/todo/v1"
)

func (s *server) AddTask(_ context.Context, in *pb.AddTaskRequest) (*pb.AddTaskResponse, error) {
	id, _ := s.d.addTask(in.Description, in.DueDate.AsTime())

	return &pb.AddTaskResponse{Id: id}, nil
}

func (s *server) ListTasks(req *pb.ListTasksRequest, stream pb.TodoService_ListTasksServer) error {
	return s.d.getTasks(func(t interface{}) error {
		task := t.(*pb.Task)
		overdue := task.DueDate != nil && !task.Done && task.DueDate.AsTime().Before(time.Now())
		err := stream.Send(&pb.ListTasksResponse{
			Task:    task,
			Overdue: overdue,
		})

		if err != nil {
			return err
		}

		return nil
	})
}

func (s *server) UpdateTasks(stream pb.TodoService_UpdateTasksServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.UpdateTasksResponse{})
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		s.d.updateTask(
			req.Task.Id,
			req.Task.Description,
			req.Task.DueDate.AsTime(),
			req.Task.Done,
		)
	}
}

func (s *server) DeleteTasks(stream pb.TodoService_DeleteTasksServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v\n", err)
		}

		s.d.deleteTask(req.Id)
		stream.Send(&pb.DeleteTasksResponse{})
	}
}
