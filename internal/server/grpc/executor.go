package grpc

import (
	"context"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
)

type ExecutorServer struct {
	serverpb.UnimplementedExecutorServiceServer
}

func (s *ExecutorServer) Sync(ctx context.Context, req *serverpb.SyncRequest) (*serverpb.SyncResponse, error) {
	panic("implement me")
}
