package main

import (
	"context"

	pb "github.com/Raghulds/go-with-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hi Dude",
	}, nil
}
