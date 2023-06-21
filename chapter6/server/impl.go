package main

import (
	"context"
	"io"
	"time"

	pb "github.com/PacktPublishing/gRPC-Go-for-Professionals/proto/todo/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"golang.org/x/exp/slices"
)

// Filter applies a mask (FieldMask) to a msg.
func Filter(msg proto.Message, mask *fieldmaskpb.FieldMask) {
	if mask == nil || len(mask.Paths) == 0 {
		return
	}

	// creates a object to apply reflection on msg
	rft := msg.ProtoReflect()

	// loop over all the fields in rft
	rft.Range(func(fd protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
		if !slices.Contains(mask.Paths, string(fd.Name())) {
			rft.Clear(fd) // clear all the fields that are not contained in mask
		}
		return true
	})
}

// AddTask adds a Task to the database.
// It returns the id of the newly inserted Task or an error.
func (s *server) AddTask(_ context.Context, in *pb.AddTaskRequest) (*pb.AddTaskResponse, error) {
	id, _ := s.d.addTask(in.Description, in.DueDate.AsTime())

	return &pb.AddTaskResponse{Id: id}, nil
}

// ListTasks streams the Tasks present in the database.
// It optionally returns an error if anything went wrong.
func (s *server) ListTasks(req *pb.ListTasksRequest, stream pb.TodoService_ListTasksServer) error {
	return s.d.getTasks(func(t interface{}) error {
		task := t.(*pb.Task)

		Filter(task, req.Mask)

		overdue := task.DueDate != nil && !task.Done && task.DueDate.AsTime().Before(time.Now().UTC())
		err := stream.Send(&pb.ListTasksResponse{
			Task:    task,
			Overdue: overdue,
		})

		return err
	})
}

// UpdateTasks apply the updates needed to be made.
// It reads the changes to be made through stream.
// It optionally returns an error if anything went wrong.
func (s *server) UpdateTasks(stream pb.TodoService_UpdateTasksServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.UpdateTasksResponse{})
		}

		if err != nil {
			return err
		}

		s.d.updateTask(
			req.Id,
			req.Description,
			req.DueDate.AsTime(),
			req.Done,
		)
	}
}

// DeleteTasks deletes Tasks in the database.
// It reads the changes to be made through stream.
// For each change being applied it sends back an acknowledgement.
// It optionally returns an error if anything went wrong.
func (s *server) DeleteTasks(stream pb.TodoService_DeleteTasksServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		s.d.deleteTask(req.Id)
		stream.Send(&pb.DeleteTasksResponse{})
	}
}
