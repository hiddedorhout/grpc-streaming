package main

import (
	"log"
	"net"

	pb ".."
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterStreamingAPIServer(grpcServer, &Routes{
		message: "Server Sided streaming",
	})
	log.Println("gRPC test server at port 8000")
	log.Fatal(grpcServer.Serve(lis))

}
