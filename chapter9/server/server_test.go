package main

import (
	"context"
	"log"
	"net"

	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/proto/todo/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
var fakeDb *FakeDb = NewFakeDb()

// init is called whenever the tests are called.
// It creates the lis (bufconn.Listener) and
// initialize the server by registering TodoService.
func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	var testServer *server = &server{
		d: fakeDb,
	}

	pb.RegisterTodoServiceServer(s, testServer)

	go func() {
		if err := s.Serve(lis); err != nil && err.Error() != "closed" {
			log.Fatalf("Server exited with error: %v\n", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
