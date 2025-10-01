package datastore

import "context"

type Node struct {
}

type NodeStore interface {
	GetNodes(ctx context.Context) ([]Node, error)
}
