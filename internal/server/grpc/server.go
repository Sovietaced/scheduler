package grpc

import (
	"context"
	"fmt"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
)

// HelloServer implements the generated SchedulerServiceServer interface.
type HelloServer struct {
	serverpb.UnimplementedServerServiceServer
}

// SayHello returns a simple greeting.
func (s *HelloServer) SayHello(ctx context.Context, req *serverpb.HelloRequest) (*serverpb.HelloReply, error) {
	name := req.GetName()
	if name == "" {
		name = "world"
	}
	return &serverpb.HelloReply{Message: fmt.Sprintf("Hello, %s!", name)}, nil
}
