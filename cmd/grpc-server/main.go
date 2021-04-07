package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/ozoncp/ocp-task-api/pkg/ocp-task-api"
)

const (
	errTaskNotFound = "task not found"

	grpcPort = ":50551"
)

type server struct {
	api.UnimplementedOcpTaskApiServer
}

func (s *server) DescribeTask(
	ctx context.Context,
	in *api.DescribeTaskRequest,
) (*api.DescribeTaskResponse, error) {

	err := status.Error(codes.NotFound, errTaskNotFound)
	return nil, err
}

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterOcpTaskApiServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
