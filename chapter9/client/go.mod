module github.com/PacktPublishing/gRPC-Go-for-Professionals/client

go 1.22

replace github.com/PacktPublishing/gRPC-Go-for-Professionals/proto => ../proto

require (
	github.com/PacktPublishing/gRPC-Go-for-Professionals/proto v0.0.0-20240314031024-bbcc94dd9932
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
)
