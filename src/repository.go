package main

import (
	"sync"
"context"
	// Import the generated protobuf code
	pb "github.com/josmel/br-seed-go/proto/consignment"
)

type repository interface {
	Lggg(ctx context.Context, *pb.LoginRequest) (*pb.LoginRequest, error)
}

type Repository struct {
	mu     sync.RWMutex
	logggs []*pb.LoginRequest
}

// Create a new consignment
func (repository *Repository) Lggg(ctx context.Context, loggg *pb.LoginRequest) (*pb.LoginRequest, error) {
	repository.mu.Lock()
	updated := append(repository.logggs, loggg)
	repository.logggs = updated
	repository.mu.Unlock()
	return loggg, nil
}
