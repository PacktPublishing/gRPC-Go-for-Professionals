module github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/server

go 1.20

replace github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/proto => ../proto

require (
	github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/proto v0.0.0-20230603071030-01673c2d8e97
	github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus v1.0.0-rc.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5
	github.com/prometheus/client_golang v1.15.1
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.42.0
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
	golang.org/x/sync v0.2.0
	golang.org/x/time v0.3.0
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	cloud.google.com/go/compute v1.19.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
