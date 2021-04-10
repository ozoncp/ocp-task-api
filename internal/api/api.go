package api

import (
	"context"

	desc "github.com/ozoncp/ocp-task-api/pkg/ocp-task-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errTaskNotFound = "task not found"
)

type api struct {
	desc.UnimplementedOcpTaskApiServer
}

func (a *api) DescribeTask(
	ctx context.Context,
	req *desc.DescribeTaskRequest,
) (*desc.DescribeTaskResponse, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := status.Error(codes.NotFound, errTaskNotFound)
	return nil, err
}

func NewOcpTaskApi() desc.OcpTaskApiServer {
	return &api{}
}
