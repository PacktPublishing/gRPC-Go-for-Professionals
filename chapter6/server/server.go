package main

import (
	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter6/proto/todo/v2"
)

type server struct {
	d db

	pb.UnimplementedTodoServiceServer
}
