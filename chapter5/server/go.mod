module github.com/PacktPublishing/gRPC-Go-for-Professionals/server

go 1.20

replace github.com/PacktPublishing/gRPC-Go-for-Professionals/proto => ../proto

require (
	github.com/PacktPublishing/gRPC-Go-for-Professionals/proto v0.0.0-20230601065932-82c46f0a6947
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
