package grpc

import (
	"context"
	"fmt"

	executorpb "github.com/sovietaced/scheduler/api/gen/pb-go/executor"
)

// HelloServer implements the generated SchedulerServiceServer interface.
type HelloServer struct {
	executorpb.UnimplementedExecutorServiceServer
}

// SayHello returns a simple greeting.
func (s *HelloServer) SayHello(ctx context.Context, req *executorpb.HelloRequest) (*executorpb.HelloReply, error) {
	name := req.GetName()
	if name == "" {
		name = "world"
	}
	return &executorpb.HelloReply{Message: fmt.Sprintf("Hello, %s!", name)}, nil
}
