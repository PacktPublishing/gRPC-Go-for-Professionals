# Chapter 8 - More essential features

This folder shows other essential features (validating, tracing, rate limiting, ...) that can be added with [protoc-gen-validate](https://github.com/bufbuild/protoc-gen-validate) and [go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware). The main code for the implementations can be found in the following files:

- [client/main.go](client/main.go)
- [client/interceptors.go](client/interceptors.go)
- [server/impl.go](server/impl.go)
- [server/interceptors.go](server/interceptors.go)
- [server/limit.go](server/limit.go)

## Running the code

### Server

#### **go run**

```shell
$ buf generate proto # or with protoc (see chapter4/README.md)
$ go run ./server 0.0.0.0:50051 0.0.0.0:50052
```

#### **bazel run**

```shell
$ bazel run //server:server 0.0.0.0:50051 0.0.0.0:50052
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

## Prometheus

Once the client made some requests to the server, Prometheus exposes metrics on the /metrics route of the HTTP server. If you ran the following command:

```shell
$ go run ./server 0.0.0.0:50051 0.0.0.0:50052
```

the second address is the HTTP server address. This means that you can access the metrics like so:

```shell
$ curl http://localhost:50052/metrics
```