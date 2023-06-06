# Chapter 9 - Production-Grade APIs

> This is based on the work done in chapter8 folder.

This folders shows how to unit and load test, and how to deploy your microservice on k8s. The main code for the implementations can be found in the following files:

- [server/server_test.go](server/server_test.go)
- [server/impl_test.go](server/impl_test.go)
- [server/Dockerfile](server/Dockerfile)
- [k8s/server.yaml](k8s/server.yaml)
- [envoy/Dockerfile](envoy/Dockerfile)
- [envoy/envoy.yaml](envoy/envoy.yaml)
- [envoy/service.yaml](envoy/service.yaml)
- [envoy/deployment.yaml](envoy/deployment.yaml)

## Running the code

### Server

```shell
$ go test ./server
```

### Docker (Server)

```shell
$ docker buildx create  --name mybuild --driver=docker-container
$ docker buildx build \
			 --tag clementjean/grpc-go-packt-book:server \
			 --file server/Dockerfile \
			 --platform linux/arm64 \ # choose your platform
		   --builder mybuild \
			 --load .
$ docker image ls
REPOSITORY                          TAG            SIZE
clementjean/grpc-go-packt-book      server         10.9MB
```

### Docker (Envoy)

```shell
$ docker buildx create  --name mybuild --driver=docker-container # if needed
$ docker buildx build \
			 --tag clementjean/grpc-go-packt-book:envoy-proxy \
			 --file server/Dockerfile \
			 --platform linux/arm64 \ # choose your platform
		   --builder mybuild \
			 --load .
REPOSITORY                          TAG            SIZE
clementjean/grpc-go-packt-book      envoy-proxy    77.1MB
```

### Kind

If you use [Kind](https://kind.sigs.k8s.io/) for your cluster, you can run the following command to spin up a 3 node cluster:

```shell
$ kind create cluster --config k8s/kind.yaml
```

### Kubernetes

```shell
$ kubectl apply -f k8s/server.yaml
$ kubectl apply -f envoy/service.yaml
$ kubectl apply -f envoy/deployment.yaml
```

If you need to port-forward, you can run:

```shell
$ kubectl get pods
NAME                READY   STATUS
todo-envoy-$HASH    1/1     Running
#...
$ kubectl port-forward pod/todo-envoy-$HASH 50051
```

Then you should be able to run the client against the port 50051.