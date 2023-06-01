//go:build !test

package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	pb "github.com/PacktPublishing/Implementing-gRPC-in-Golang-Microservice/proto/todo/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	//_ "google.golang.org/grpc/encoding/gzip"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

func newMetricsServer(httpAddr string, reg *prometheus.Registry) *http.Server {
	httpSrv := &http.Server{Addr: httpAddr}
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	httpSrv.Handler = m
	return httpSrv
}

func newGrpcServer(lis net.Listener, srvMetrics *grpcprom.ServerMetrics) (*grpc.Server, error) {
	creds, err := credentials.NewServerTLSFromFile("./certs/server_cert.pem", "./certs/server_key.pem")
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime)
	// limiter := &simpleLimiter{
	// 	limiter: rate.NewLimiter(2, 4),
	// }

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			//ratelimit.UnaryServerInterceptor(limiter),
			otelgrpc.UnaryServerInterceptor(),
			srvMetrics.UnaryServerInterceptor(),
			auth.UnaryServerInterceptor(validateAuthToken),
			logging.UnaryServerInterceptor(logCalls(logger)),
		),
		grpc.ChainStreamInterceptor(
			//ratelimit.StreamServerInterceptor(limiter),
			otelgrpc.StreamServerInterceptor(),
			srvMetrics.StreamServerInterceptor(),
			auth.StreamServerInterceptor(validateAuthToken),
			logging.StreamServerInterceptor(logCalls(logger)),
		),
	}
	s := grpc.NewServer(opts...)

	pb.RegisterTodoServiceServer(s, &server{
		d: New(),
	})
	reflection.Register(s)

	return s, nil
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		log.Fatalln("usage: server [GRPC_IP_ADDR] [METRICS_IP_ADDR]")
	}

	grpcAddr := args[0]
	httpAddr := args[1]

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// handle CTRL+C
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(quit)

	g, ctx := errgroup.WithContext(ctx)

	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg := prometheus.NewRegistry()
	reg.MustRegister(srvMetrics)

	lis, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	grpcServer, err := newGrpcServer(lis, srvMetrics)

	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	g.Go(func() error {
		log.Printf("gRPC server listening at %s\n", grpcAddr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to gRPC server: %v\n", err)
			return err
		}
		log.Println("gRPC server shutdown")
		return nil
	})

	metricsServer := newMetricsServer(httpAddr, reg)
	g.Go(func() error {
		log.Printf("metrics server listening at %s\n", httpAddr)
		if err := metricsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to serve metrics: %v\n", err)
			return err
		}
		log.Println("metrics server shutdown")
		return nil
	})

	// handle termination
	select {
	case <-quit:
		break
	case <-ctx.Done():
		break
	}

	// gracefully shutdown servers
	cancel()

	timeoutCtx, timeoutCancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer timeoutCancel()

	log.Println("shutting down servers, please wait...")

	grpcServer.GracefulStop()
	metricsServer.Shutdown(timeoutCtx)

	// wait for shutdown
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
