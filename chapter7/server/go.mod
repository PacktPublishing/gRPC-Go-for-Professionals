module github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter7/server

go 1.20

replace github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter7/proto => ../proto

require (
	github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter7/proto v0.0.0-20230526092912-af064518b681
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230526203410-71b5a4ffd15e // indirect
)
