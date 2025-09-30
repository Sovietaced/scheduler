package grpc

import (
	"context"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
)

type WorkloadServer struct {
	serverpb.UnimplementedWorkloadServiceServer
}

func (w *WorkloadServer) CreateWorkload(ctx context.Context, req *serverpb.CreateWorkloadRequest) (*serverpb.CreateWorkloadResponse, error) {
	return &serverpb.CreateWorkloadResponse{}, nil
}
