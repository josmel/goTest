package main

import (
	"context"

	// Import the generated protobuf code
	pb "github.com/josmel/br-seed-go/proto/consignment"
)

type handler struct {
	repository
}

func (s *handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {

	// Save our consignment
	data, err := s.repository.Lggg(ctx, req)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, err
	}

	return &pb.LoginReply{ApiConnectToken: "pepe", WasCookie: "dd"}, nil
}
