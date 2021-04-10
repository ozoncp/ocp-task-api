package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-task-api/pkg/ocp-task-api"
)

const (
	errTaskNotFound = "task not found"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:82", "gRPC server endpoint")
	restServerEndpoint = "0.0.0.0:80"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterOcpTaskApiHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	fmt.Printf("Server listening on %s\n", restServerEndpoint)
	return http.ListenAndServe(restServerEndpoint, mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
