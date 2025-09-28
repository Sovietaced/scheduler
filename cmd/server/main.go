package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
	servergrpc "github.com/sovietaced/scheduler/internal/server/grpc"
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
	// Enable server reflection for easier debugging with tools like grpcurl.
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server exited: %v", err)
	}
}
