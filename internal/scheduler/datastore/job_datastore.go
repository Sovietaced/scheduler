package datastore

import "context"

type Job struct {
}

type JobDatastore interface {
	// Get non-terminal jobs
	GetJobsSnapshot(ctx context.Context) ([]Job, error)
}
