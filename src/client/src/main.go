package main

import (
	"context"
	"io"
	"log"
	"time"

	pb ".."
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := pb.NewStreamingAPIClient(conn)
	msg, err := client.DisplayMessage(ctx, &pb.DisplayRequest{
		ProcessID: "0000",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(msg.GetInfo())

	stream, err := client.DisplayStream(ctx, &pb.DisplayRequest{
		ProcessID: "0000",
	})
	if err != nil {
		log.Println(err)
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Println(err)
		}
		if err == io.EOF {
			break
		}
		log.Println(msg.GetInfo())
	}

}
