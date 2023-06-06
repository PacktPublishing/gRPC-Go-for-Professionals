# Chapter 7 - Out of the box features

> This is based on the work done in chapter6 folder.

This folder shows how more advanced features that comes out of the box with gRPC (error handling, interceptors, load balancing, ...). The main code for the implementations can be found in the following files:

- [client/main.go](client/main.go)
- [client/interceptors.go](client/interceptors.go)
- [server/impl.go](server/impl.go)
- [server/interceptors.go](server/interceptors.go)

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

### Kubernetes (k8s)

In order to show the load balancing effects, this folder introduces configurations for deploying the server and client to Kubernetes.

> A Kubernetes is required in order to make the following command work. If you have [kind](https://kind.sigs.k8s.io/) installed, you can run the following command to quickly spin up a cluster: `kind create cluster --config k8s/kind.yaml`.

#### **Server**

```shell
$ kubectl apply -f k8s/server.yaml
```

#### **Client**

```shell
$ kubectl apply -f k8s/client.yaml
```