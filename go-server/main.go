package main

import (
	"fmt"
	"log"
	"net"

	posts "github.com/leomirandadev/grpc-project/go-server/posts"
	"google.golang.org/grpc"
)

var (
	serverURL  = "localhost:10000"
	portServer = 10000
)

func main() {

	fmt.Println("Server running on", serverURL)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", portServer))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := posts.Server{}

	grpcServer := grpc.NewServer()

	posts.RegisterPostServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
