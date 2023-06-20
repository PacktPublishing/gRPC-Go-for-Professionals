package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/helpers/proto"
)

// serializedSize calculates the serialized size of the msg.
// It returns the number of bytes in which the msg was serialized.
func serializedSize[M protoreflect.ProtoMessage](msg M) int {
	out, err := proto.Marshal(msg)

	if err != nil {
		log.Fatal(err)
	}

	return len(out)
}

func main() {
	t := &pb.Tags{}
	tags := []int{1, 16, 2048, 262_144, 33_554_432, 536_870_911}
	fields := []*int32{&t.Tag, &t.Tag2, &t.Tag3, &t.Tag4, &t.Tag5, &t.Tag6}

	sz := serializedSize(t)
	fmt.Printf("0 - %d\n", sz)

	for i, f := range fields {
		*f = 1

		sz := serializedSize(t)
		fmt.Printf("%d - %d\n", tags[i], sz-(i+1))
	}
}
