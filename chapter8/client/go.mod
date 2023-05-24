module github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter8/client

go 1.20

replace github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter8/proto => ../proto

require (
	github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter8/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)
