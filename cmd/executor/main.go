package main

import (
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	executorpb "github.com/sovietaced/scheduler/api/gen/pb-go/executor"
	executorgrpc "github.com/sovietaced/scheduler/internal/executor/grpc"
)

func main() {
	port := "8081"
	addr := ":" + port

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", addr, err)
	}
	log.Printf("gRPC server listening on %s", addr)

	grpcServer := grpc.NewServer()
	// Register our hello world service.
	executorpb.RegisterExecutorServiceServer(grpcServer, &executorgrpc.HelloServer{})
	// Enable server reflection for easier debugging with tools like grpcurl.
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server exited: %v", err)
	}
}
