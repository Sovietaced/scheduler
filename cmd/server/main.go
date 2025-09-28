package main

import (
	"context"
	"log"
	"net"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
	"github.com/sovietaced/scheduler/internal/scheduler"
	servergrpc "github.com/sovietaced/scheduler/internal/server/grpc"
	"google.golang.org/grpc"
)

func main() {
	port := "8080"
	addr := ":" + port

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", addr, err)
	}
	log.Printf("gRPC server listening on %s", addr)

	grpcServer := grpc.NewServer()

	serverpb.RegisterQueueServiceServer(grpcServer, &servergrpc.QueueServer{})
	serverpb.RegisterExecutorServiceServer(grpcServer, &servergrpc.ExecutorServer{})

	s := scheduler.NewScheduler()
	go func() {
		s.Run(context.Background())
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server exited: %v", err)
	}
}
