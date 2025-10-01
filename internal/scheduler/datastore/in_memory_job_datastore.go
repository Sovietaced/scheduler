package datastore

import "context"

type InMemoryJobDataStore struct {
	jobs []Job
}

func NewInMemoryJobDataStore() *InMemoryJobDataStore {
	return &InMemoryJobDataStore{jobs: make([]Job, 0)}
}

func (i *InMemoryJobDataStore) GetJobsSnapshot(ctx context.Context) ([]Job, error) {
	return i.jobs, nil //FIXME:  return a copy
}
