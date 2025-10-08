package executor

import (
	"context"
	"fmt"
	"time"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
	"k8s.io/client-go/kubernetes"
)

type Executor struct {
	executorClient      serverpb.ExecutorServiceClient
	clusterStateManager *ClusterStateManager
}

func NewExecutor(ctx context.Context, executorClient serverpb.ExecutorServiceClient, kubernetesClient kubernetes.Interface) *Executor {
	return &Executor{executorClient: executorClient, clusterStateManager: NewClusterStateManager(ctx, kubernetesClient)}
}

func (e *Executor) Run(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("syncing")
			_, err := e.executorClient.Sync(ctx, &serverpb.SyncRequest{})
			if err != nil {
				fmt.Println(fmt.Sprintf("failed to sync: %v", err))
			}

			_, err = e.clusterStateManager.GetClusterState()
			if err != nil {
				fmt.Println(fmt.Sprintf("cluster state manager error: %v", err))
			}

		}

	}
}
