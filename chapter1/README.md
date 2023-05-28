# Chapter 1 - Networking Primer

The files contained in this folder are [Wireshark](https://www.wireshark.org/) capture files. Each of them show the logs for different types of gRPC API. Their purpose is to illustrate the HTTP/2 protocol used by gRPC. You can view the captures in Wireshark by opening them (`File > Open`), and adding the following display filter:

```
tcp.port == 50051 and (grpc or http2)
```