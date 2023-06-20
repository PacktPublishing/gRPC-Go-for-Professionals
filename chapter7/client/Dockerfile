# This is supposed to be build with the BuildKit engine
#
# Params:
# SERVER_ADDR - this is the address of the server that will be passed as comand line arg (default: 0.0.0.0:50051)
#
# e.g. to build for linux/arm64 (can also be linux/amd64), load it in Docker, and for k8s demo
# sudo docker buildx create  --name mybuild --driver=docker-container
# sudo docker buildx build \
# 		--tag clementjean/grpc-go-packt-book:client \
# 		--file client/Dockerfile \
# 		--build-arg SERVER_ADDR="dns:///todo-server.default.svc.cluster.local:50051" \
# 		--platform linux/arm64 \
# 		--builder mybuild \
# 		--load .

FROM --platform=$BUILDPLATFORM alpine:latest as protoc
ARG BUILDPLATFORM TARGETARCH

# download the protoc binary from github
RUN export PROTOC_VERSION=23.2 \
    && export PROTOC_ARCH=$(uname -m | sed s/aarch64/aarch_64/) \
    && export PROTOC_OS="linux" \
    && export PROTOC_ZIP=protoc-$PROTOC_VERSION-$PROTOC_OS-$PROTOC_ARCH.zip \
    && echo "downloading: " https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_ZIP \
    && wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc 'include/*' \
    && rm -f $PROTOC_ZIP

FROM --platform=$BUILDPLATFORM golang:1.20-alpine as build
ARG BUILDPLATFORM TARGETARCH

# copy the protoc binary and the protobuf includes
COPY --from=protoc /usr/local/bin/protoc /usr/local/bin/protoc
COPY --from=protoc /usr/local/include/google /usr/local/include/google

# download protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# copy proto into go/src/proto
WORKDIR /go/src/proto
COPY ./proto .

# copy code into go/src/app
WORKDIR /go/src/app
COPY ./client .

# download dependencies and build
RUN go mod download
RUN CGO_ENABLED=0 GOOS="linux" GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/bin/app

# here we use an alpine image because if we were using a scratch image (like in the server image)
# we would not be able to pass SERVER_ADDR_ENV variable as command line arg.
# this make the image a bit bigger.
FROM alpine:latest
ARG SERVER_ADDR="0.0.0.0:50051"

# copy certs into /certs
COPY ./certs/ca_cert.pem ./certs/ca_cert.pem

# copy the previously built binary into smaller image
COPY --from=build /go/bin/app /
ENV SERVER_ADDR_ENV $SERVER_ADDR
CMD /app ${SERVER_ADDR_ENV}