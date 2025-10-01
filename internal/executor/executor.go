package executor

import (
	"context"
	"fmt"
	"time"

	serverpb "github.com/sovietaced/scheduler/api/gen/pb-go/server"
)

type Executor struct {
	executorClient serverpb.ExecutorServiceClient
}

func NewExecutor(executorClient serverpb.ExecutorServiceClient) *Executor {
	return &Executor{executorClient: executorClient}
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
				//FIXME: do something
			}
		}

	}
}
