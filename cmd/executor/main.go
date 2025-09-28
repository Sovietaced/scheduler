package main

import (
	"context"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
	"github.com/sovietaced/scheduler/internal/executor"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := serverpb.NewExecutorServiceClient(conn)
	ex := executor.NewExecutor(client)
	ex.Run(context.Background())
}
