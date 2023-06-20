# This is supposed to be build with the BuildKit engine
#
# e.g. to build for linux/arm64 (can also be linux/amd64) and load it in Docker
# sudo docker buildx create  --name mybuild --driver=docker-container
# sudo docker buildx build \
# 	--tag clementjean/grpc-go-packt-book:server \
# 	--file server/Dockerfile \
# 	--platform linux/arm64 \
# 	--builder mybuild \
# 	--load .

FROM --platform=$BUILDPLATFORM alpine as protoc
ARG BUILDPLATFORM TARGETOS TARGETARCH

# download the protoc binary from github
RUN export PROTOC_VERSION=23.2 \
    && export PROTOC_ARCH=$(uname -m | sed s/aarch64/aarch_64/) \
    && export PROTOC_OS=$(echo $TARGETOS | sed s/darwin/linux/) \
    && export PROTOC_ZIP=protoc-$PROTOC_VERSION-$PROTOC_OS-$PROTOC_ARCH.zip \
    && echo "downloading: " https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_ZIP \
    && wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc 'include/*' \
    && rm -f $PROTOC_ZIP

FROM --platform=$BUILDPLATFORM golang:1.20-alpine as build
ARG BUILDPLATFORM TARGETOS TARGETARCH

# copy the protoc binary and the protobuf includes
COPY --from=protoc /usr/local/bin/protoc /usr/local/bin/protoc
COPY --from=protoc /usr/local/include/google /usr/local/include/google

# download protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/envoyproxy/protoc-gen-validate@latest

# copy proto files into go/src/proto
WORKDIR /go/src/proto
COPY ./proto .

# generate code out of proto files
RUN protoc -I. \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --validate_out="lang=go,paths=source_relative:." \
    **/*.proto

# copy code into go/src/server
WORKDIR /go/src/server
COPY ./server .

# download dependencies and build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/bin/server

FROM scratch

# copy certs into /certs
COPY ./certs/server_cert.pem ./certs/server_cert.pem
COPY ./certs/server_key.pem ./certs/server_key.pem

# copy the previously built binary into smaller image
COPY --from=build /go/bin/server /
EXPOSE 50051 50052
CMD ["/server", "0.0.0.0:50051", "0.0.0.0:50052"]