# Chapter 5 - Type of gRPC Endpoints

> This is based on the work done in chapter4 folder.

This folder shows how to implement the different types of gRPC APIs (unary, server streaming, client streaming, and bidirectional streaming). The main code for the implementations can be found in the following files:

- [client/main.go](client/main.go)
- [server/impl.go](server/impl.go)

## Running the code

### Server

#### **go run**

```shell
$ buf generate proto # or with protoc (see chapter4/README.md)
$ go run ./server 0.0.0.0:50051
```

#### **bazel run**

```shell
$ bazel run //server:server 0.0.0.0:50051
```

### Client

#### **go run**

```shell
# if not done before `buf generate proto` or with protoc (see chapter4/README.md)
$ go run ./client 0.0.0.0:50051
```

#### **bazel run**

```shell
$ bazel run //client:client 0.0.0.0:50051
```