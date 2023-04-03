package main

import (
	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter5/proto/todo/v1"
)

type server struct {
	d db

	pb.UnimplementedTodoServiceServer
}
