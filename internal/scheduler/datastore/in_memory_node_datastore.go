package datastore

import "context"

type InMemoryNodeDataStore struct {
	nodes []Node
}

func NewInMemoryNodeDataStore() *InMemoryNodeDataStore {
	return &InMemoryNodeDataStore{nodes: make([]Node, 0)}
}

func (s *InMemoryNodeDataStore) GetNodes(ctx context.Context) ([]Node, error) {
	return s.nodes, nil //FIXME: Return a copy?
}
