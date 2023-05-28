# Chapter 6 - APIs’ Design Considerations

This folder shows how to improve the API presented in the chapter5 folder. The main code for the implementations can be found in the following files:

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