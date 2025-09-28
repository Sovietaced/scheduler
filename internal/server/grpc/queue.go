package grpc

import (
	"context"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
)

type QueueServer struct {
	serverpb.UnimplementedQueueServiceServer
}

func (q *QueueServer) Sync(ctx context.Context, request *serverpb.SyncRequest) (*serverpb.SyncResponse, error) {
	//TODO implement me
	panic("implement me")
}
