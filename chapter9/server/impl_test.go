package main

import (
	"context"
	"io"
	"testing"
	"time"

	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/proto/todo/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	errorNoDatabaseAccess   = "unexpected error: couldn't access the database"
	errorInvalidDescription = "invalid AddTaskRequest.Description: value length must be at least 1 runes"
	errorInvalidDueDate     = "invalid AddTaskRequest.DueDate: value must be greater than now"
)

// newClient creates a new client for tests.
// It uses the bufDialer func (see server/server_test.go)
// which lets the created client to talk to the server
// through a in memory channel.
func newClient(t *testing.T) (*grpc.ClientConn, pb.TodoServiceClient) {
	ctx := context.TODO()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}

	return conn, pb.NewTodoServiceClient(conn)
}

// errorIs checks that err has the given code and error message
// this is only useful for gRPC errors (it uses status.FromError).
// It returns wether the error is 'as expected'.
func errorIs(err error, code codes.Code, msg string) bool {
	if err != nil {
		if s, ok := status.FromError(err); ok {
			if code == s.Code() && s.Message() == msg {
				return true
			}
		}
	}

	return false
}

func TestRunAll(t *testing.T) {
	t.Run("AddTaskTests", func(t *testing.T) {
		t.Run("TestAddTaskUnavailableDb", testAddTaskUnavailableDb)
		t.Run("TestAddTaskEmptyDescription", testAddTaskEmptyDescription)
		t.Run("TestAddTaskWrongDueDate", testAddTaskWrongDueDate)
		t.Run("TestAddTask", testAddTask)
	})

	t.Run("ListTasksTests", func(t *testing.T) {
		t.Run("TestListTasksUnavailableDb", testListTasksUnavailableDb)
		t.Run("TestListTasks", testListTasks)
	})

	t.Run("UpdateTasksTests", func(t *testing.T) {
		t.Run("TestUpdateTasksUnavailableDb", testUpdateTasksUnavailableDb)
		t.Run("TestUpdateTasks", testUpdateTasks)
	})

	t.Run("DeleteTasksTests", func(t *testing.T) {
		t.Run("TestDeleteTasksUnavailableDb", testDeleteTasksUnavailableDb)
		t.Run("TestDeleteTasks", testDeleteTasks)
	})

	t.Cleanup(func() {
		lis.Close()
	})
}

// Tests
func testAddTaskUnavailableDb(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	newDb := NewFakeDb(IsAvailable(false))
	*fakeDb = *newDb

	req := &pb.AddTaskRequest{
		Description: "test",
		DueDate:     timestamppb.New(time.Now().Add(5 * time.Hour)),
	}

	// Act
	_, err := c.AddTask(context.TODO(), req)
	fakeDb.Reset()

	// Assert
	if !errorIs(err, codes.Internal, errorNoDatabaseAccess) {
		t.Errorf("expected Internal, got %v", err)
	}
}

func testAddTaskEmptyDescription(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	req := &pb.AddTaskRequest{}

	// Act
	_, err := c.AddTask(context.TODO(), req)

	// Assert
	if !errorIs(err, codes.Unknown, errorInvalidDescription) {
		t.Errorf(
			"expected Unknown with message \"%s\", got %v",
			errorInvalidDescription, err,
		)
	}
}

func testAddTaskWrongDueDate(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	fakeDb.Reset()

	wrongDueDates := []*timestamppb.Timestamp{
		&timestamppb.Timestamp{},
		timestamppb.New(time.Now().UTC()),
		timestamppb.New(time.Now().UTC().Add(-1 * time.Hour)),
	}

	for i, wrongDueDate := range wrongDueDates {
		req := &pb.AddTaskRequest{
			Description: "test",
			DueDate:     wrongDueDate,
		}

		// Act
		_, err := c.AddTask(context.TODO(), req)

		// Assert
		if err == nil || !errorIs(err, codes.Unknown, errorInvalidDueDate) {
			t.Errorf(
				"test %d, expected Unknown with message \"%s\", got %v",
				i, errorInvalidDueDate, err,
			)
		}
	}
}

func testAddTask(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	req := &pb.AddTaskRequest{
		Description: "test",
		DueDate:     timestamppb.New(time.Now().Add(5 * time.Hour)),
	}

	// Act
	_, err := c.AddTask(context.TODO(), req)

	// Assert
	if err != nil {
		t.Errorf(
			"expected error nil, got %v",
			err,
		)
	}
}

func testListTasksUnavailableDb(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	newDb := NewFakeDb(IsAvailable(false))
	*fakeDb = *newDb

	req := &pb.ListTasksRequest{}

	// Act
	var errRecv error
	res, err := c.ListTasks(context.TODO(), req)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if res != nil {
		_, err = res.Recv()
	}

	fakeDb.Reset()

	// Assert
	if !errorIs(err, codes.Unknown, errorNoDatabaseAccess) {
		t.Errorf("expected Unknown, got %v", errRecv)
	}
}

func testListTasks(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	fakeDb.d.tasks = []*pb.Task{
		{}, {}, {}, // 3 empty tasks
	}
	expectedRead := len(fakeDb.d.tasks)

	req := &pb.ListTasksRequest{}
	count := 0

	// Act
	res, err := c.ListTasks(context.TODO(), req)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	for {
		_, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Errorf("error while reading stream: %v", err)
		}

		count++
	}

	// Assert
	if count != expectedRead {
		t.Errorf(
			"expected reading %d tasks, read %d",
			expectedRead, count,
		)
	}
}

func testUpdateTasksUnavailableDb(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	newDb := NewFakeDb(IsAvailable(false))
	*fakeDb = *newDb

	// Clear all the fields in fakeDb.d.tasks
	// except Id.
	requests := []*pb.UpdateTasksRequest{
		{Id: 0}, {Id: 1}, {Id: 2},
	}

	// Act
	stream, err := c.UpdateTasks(context.TODO())

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	for _, req := range requests {
		stream.Send(req)
	}

	_, err = stream.CloseAndRecv()
	fakeDb.Reset()

	// Assert
	if !errorIs(err, codes.Internal, errorNoDatabaseAccess) {
		t.Errorf("expected Internal, got %v", err)
	}
}

func testUpdateTasks(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	fakeDb.d.tasks = []*pb.Task{
		{Id: 0, Description: "test1"},
		{Id: 1, Description: "test2"},
		{Id: 2, Description: "test3"},
	}

	// Clear all the fields in fakeDb.d.tasks
	// except Id.
	requests := []*pb.UpdateTasksRequest{
		{Id: 0}, {Id: 1}, {Id: 2},
	}
	expectedUpdates := len(requests)

	// Act
	stream, err := c.UpdateTasks(context.TODO())
	count := 0

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	for _, req := range requests {
		if err := stream.Send(req); err != nil {
			t.Fatal(err)
		}

		count++
	}

	_, err = stream.CloseAndRecv()

	// Assert
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if count != expectedUpdates {
		t.Errorf(
			"expected updating %d tasks, updated %d",
			expectedUpdates, count,
		)
	}
}

type countAndError struct {
	count int
	err   error
}

func sendRequestsOverStream(stream pb.TodoService_DeleteTasksClient, requests []*pb.DeleteTasksRequest, waitc chan countAndError) {
	for _, req := range requests {
		if err := stream.Send(req); err != nil {
			waitc <- countAndError{err: err}
			close(waitc)
			return
		}
	}

	if err := stream.CloseSend(); err != nil {
		waitc <- countAndError{err: err}
		close(waitc)
	}
}

func readResponsesOverStream(stream pb.TodoService_DeleteTasksClient, waitc chan countAndError) {
	count := 0

	for {
		_, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			waitc <- countAndError{err: err}
			close(waitc)
			return
		}

		count++
	}

	waitc <- countAndError{count: count}
	close(waitc)
}

func testDeleteTasksUnavailableDb(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	newDb := NewFakeDb(IsAvailable(false))
	*fakeDb = *newDb

	waitc := make(chan countAndError)
	requests := []*pb.DeleteTasksRequest{
		{Id: 1}, {Id: 2}, {Id: 3},
	}

	// Act
	stream, err := c.DeleteTasks(context.TODO())

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	go sendRequestsOverStream(stream, requests, waitc)
	go readResponsesOverStream(stream, waitc)

	countAndError := <-waitc
	fakeDb.Reset()

	// Assert
	if !errorIs(countAndError.err, codes.Internal, errorNoDatabaseAccess) {
		t.Errorf("expected Internal, got %v", countAndError.err)
	}
}

func testDeleteTasks(t *testing.T) {
	// Arrange
	conn, c := newClient(t)
	defer conn.Close()

	fakeDb.d.tasks = []*pb.Task{
		{Id: 1}, {Id: 2}, {Id: 3},
	}
	expectedRead := len(fakeDb.d.tasks)

	waitc := make(chan countAndError)
	requests := []*pb.DeleteTasksRequest{
		{Id: 1}, {Id: 2}, {Id: 3},
	}

	// Act
	stream, err := c.DeleteTasks(context.TODO())

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	go sendRequestsOverStream(stream, requests, waitc)
	go readResponsesOverStream(stream, waitc)

	countAndError := <-waitc

	// Assert
	if countAndError.err != nil {
		t.Errorf("expected error: %v", countAndError.err)
	}
	if countAndError.count != expectedRead {
		t.Errorf(
			"expected reading %d responses, read %d",
			expectedRead, countAndError.count,
		)
	}
}
