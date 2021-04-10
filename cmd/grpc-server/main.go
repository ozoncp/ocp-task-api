package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-task-api/internal/api"
	desc "github.com/ozoncp/ocp-task-api/pkg/ocp-task-api"
)

const (
	grpcPort = ":82"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpTaskApiServer(s, api.NewOcpTaskApi())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
