package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb ".."
)

type Routes struct {
	message string
}

func (s *Routes) DisplayMessage(ctx context.Context, request *pb.DisplayRequest) (*pb.DisplayResponse, error) {
	log.Println(request.GetProcessID())
	return &pb.DisplayResponse{
		Code: "1",
		Info: s.message,
	}, nil
}

func (s *Routes) DisplayStream(request *pb.DisplayRequest, stream pb.StreamingAPI_DisplayStreamServer) error {
	log.Println(request.GetProcessID())
	for _, step := range []int{0, 1, 2} {
		stream.Send(&pb.DisplayResponse{
			Code: "2",
			Info: fmt.Sprintf("%s: %d", s.message, step),
		})
		time.Sleep(2 * time.Second)
	}

	return nil
}
