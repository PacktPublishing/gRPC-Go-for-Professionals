# Chapter 3 - Introduction to gRPC

This folder contains the result of comparing Protobuf and JSON in order to explain one of the reasons gRPC is faster than traditional REST APIs. You can find the result of the size comparison by taking a look at the files named `accounts.bin.gz` and `accounts.json.gz`. And you can execute the serialization/deserialization time comparison by running the `main.go` file.

## Size

### Linux/Mac

```shell
$ ls -lh *.gz
571K accounts.bin.gz
650K accounts.json.gz
```

### Windows (Powershell)

```shell
$ Get-ChildItem *.gz | Select-Object Name, @{Name="Size (KB)";Expression={$_.Length / 1KB}}
Name             Size (KB)
----             ---------
accounts.bin.gz    570.806
accounts.json.gz   650.103
```

## Serialization/Deserialization Time

> Be aware that if you have generated gRPC code (proto/account_grpc.pb.go) before you will have a dependency error. Simply remove proto/account_grpc.pb.go before running the following commands.

```shell
$ gzip -dk *.gz
$ protoc --proto_path=proto \
  --go_out=. --go_opt=module=github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter3 \
  proto/*.proto
$ go run main.go
JSON: 40.520000ms
PB: 9.450000ms
```

## Generating both Protobuf and gRPC code

```shell
$ protoc --proto_path=proto \
  --go_out=. --go_opt=module=github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter3 \
  --go-grpc_out=. --go-grpc_opt=module=github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/chapter3 \
  proto/*.proto
```
