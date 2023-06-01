# Chapter 4 - Setting up a Project

This folder is a template for gRPC project that is used throughout the remaining chapters. It uses a go workspace and multiple modules (proto, server, client). This folder also shows how to set different ways of building a gRPC project.

## protoc

```shell
$ protoc --go_out=. --go_opt=module=github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice \
  --go-grpc_out=. --go-grpc_opt=module=github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice \
  proto/dummy/v1/dummy.proto
$ go run ./server 0.0.0.0:50051
listening at 0.0.0.0:50051
```

## Buf

```shell
$ buf generate proto
$ go run ./server 0.0.0.0:50051
listening at 0.0.0.0:50051
```

## Bazel

```shell
$ bazel run //:server
listening at 0.0.0.0:50051
```
