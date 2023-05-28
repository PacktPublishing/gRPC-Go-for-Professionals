package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/helpers/proto"
)

// compressedSize calculates the number of bytes after compression with gzip.
// It returns two values. The number of bytes the msg takes after serialization
// and the number of bytes after compression.
func compressedSize[M protoreflect.ProtoMessage](msg M) (int, int) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	out, err := proto.Marshal(msg)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := gz.Write(out); err != nil {
		log.Fatal(err)
	}

	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}

	return len(out), len(b.Bytes())
}

func main() {
	// var data int32 = 268435456
	// i32 := &wrapperspb.Int32Value{
	// 	Value: data,
	// }

	task := &pb.Task{
		Id: 1,
		Description: `This is a task that is quite long and requires a lot of work.
		We are not sure we can finish it even after 5 days.
		Some planning will be needed and a meeting is required.`,
		DueDate: timestamppb.New(time.Now().Add(5 * 24 * time.Hour)),
	}

	o, c := compressedSize(task)
	fmt.Printf("original: %d\ncompressed: %d\n", o, c)
}
