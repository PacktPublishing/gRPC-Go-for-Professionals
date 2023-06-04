FROM envoyproxy/envoy-distroless:v1.26-latest
COPY ./envoy/envoy.yaml /etc/envoy/envoy.yaml
COPY ./certs/server_cert.pem /etc/envoy/certs/server_cert.pem
COPY ./certs/server_key.pem /etc/envoy/certs/server_key.pem
EXPOSE 50051
CMD ["--config-path", "/etc/envoy/envoy.yaml"]