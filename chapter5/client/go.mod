module github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter5/client

go 1.20

replace github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter5/proto => ../proto

require (
	github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter5/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230327215041-6ac7f18bb9d5 // indirect
)
