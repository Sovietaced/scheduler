package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/sovietaced/scheduler/internal/scheduler/datastore"
)

type Scheduler struct {
	jobStore  datastore.JobDatastore
	nodeStore datastore.NodeStore
}

func NewScheduler(jobStore datastore.JobDatastore, nodeStore datastore.NodeStore) *Scheduler {
	return &Scheduler{jobStore: jobStore, nodeStore: nodeStore}
}

func (s *Scheduler) Run(ctx context.Context) {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.schedule(ctx)
		}

	}
}

func (s *Scheduler) schedule(ctx context.Context) {
	fmt.Println("Scheduler is running")

	// read state form db
	jobs, err := s.jobStore.GetJobsSnapshot(ctx)
	if err != nil {
		//FIXME:  do something
	}

	// schedule
	// iterate over the jobs and try and fit them onto available nodes that match

	nodes, err := s.nodeStore.GetNodes(ctx)
	if err != nil {
		//FIXME: do something
	}

	fmt.Println(nodes)

	for _, job := range jobs {
		fmt.Println(job)
	}

	// Write out events?

}
