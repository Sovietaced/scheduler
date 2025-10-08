package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
	"github.com/sovietaced/scheduler/internal/executor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Create a cancellable context that is canceled on SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(fmt.Sprintf("unable to get k8s cluster config: %v", err))
	}

	kubernetesClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("unable to create k8s client: %v", err))
	}

	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := serverpb.NewExecutorServiceClient(conn)
	ex := executor.NewExecutor(ctx, client, kubernetesClient)
	ex.Run(ctx)
}
