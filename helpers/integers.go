package main

import (
	"fmt"
	"log"
	"unsafe"

	"golang.org/x/exp/constraints"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// serializedSize calculates the size in memory of data
// and the size after Protobuf serialization of the wrapper.
// It returns two values. The number of bytes for data in memory
// and the number of bytes after serialization of wrapper - 1 (removes the byte for tag + wire type).
func serializedSize[D constraints.Integer, W protoreflect.ProtoMessage](data D, wrapper W) (uintptr, int) {
	out, err := proto.Marshal(wrapper)

	if err != nil {
		log.Fatal(err)
	}

	return unsafe.Sizeof(data), len(out) - 1
}

func main() {
	var data int32 = 268435456
	i32 := &wrapperspb.Int32Value{
		Value: data,
	}

	d, w := serializedSize(data, i32)
	fmt.Printf("in memory: %d\npb: %d\n", d, w)
}
