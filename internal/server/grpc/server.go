package grpc

import (
	"context"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
)

type QueueServer struct {
	serverpb.UnimplementedQueueServiceServer
}

func (q *QueueServer) CreateQueue(ctx context.Context, req *serverpb.CreateQueueRequest) (*serverpb.CreateQueueResponse, error) {
	panic("implement me")
}
