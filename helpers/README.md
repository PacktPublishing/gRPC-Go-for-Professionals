# Helpers

## Proto

In order to build the proto files in the `proto` folder, you can run the following command:

```shell
protoc -Iproto --go_out=proto --go_opt=paths=source_relative proto/*.proto
```

This should generate files and give you the following:

```shell
proto
├── split.pb.go
├── split.proto
├── tags.pb.go
├── tags.proto
├── todo.pb.go
└── todo.proto
```

## Go

> Before you read this section, make sure that you generated the go code from the proto files (see previous section).

There are multiple `main` functions in this folder so you cannot do a simple `go run`. You will have to specify the helper you want to run:

### Integers

```shell
$ go run integers.go
in memory: 4
pb: 5
```

### Tags

```shell
$ go run tags.go
0 - 0
1 - 1
16 - 3
2048 - 6
262144 - 10
33554432 - 15
536870911 - 20
```

### Split

```shell
$ go run split.go
With Name: 7
With ComplexName: 9
```

### Gzip

```shell
$ go run gzip.go
original: 191
compressed: 180
```