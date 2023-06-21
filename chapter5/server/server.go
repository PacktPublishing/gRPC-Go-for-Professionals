package main

import (
	pb "github.com/PacktPublishing/gRPC-Go-for-Professionals/proto/todo/v1"
)

type server struct {
	d db

	pb.UnimplementedTodoServiceServer
}
