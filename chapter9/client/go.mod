module github.com/PacktPublishing/gRPC-Go-for-Professionals/client

go 1.20

replace github.com/PacktPublishing/gRPC-Go-for-Professionals/proto => ../proto

require (
	github.com/PacktPublishing/gRPC-Go-for-Professionals/proto v0.0.0-20230603071030-01673c2d8e97
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/stretchr/testify v1.8.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
