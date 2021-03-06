package main

import (
	"context"
	"fmt"
	"log"
	"time"

	services "github.com/leomirandadev/grpc-project/go-client/services"
	"google.golang.org/grpc"
)

var (
	serverURL = "localhost:10000"
)

func getGRPCClient() *grpc.ClientConn {
	var opts = []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.Dial(serverURL, opts...)

	if err != nil {
		log.Fatal("Fail to dial %v", err)
	}

	return conn
}

func main() {
	conn := getGRPCClient()

	defer conn.Close()

	client := services.NewPostServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	posts, err := client.GetPosts(ctx, &services.Empty{})

	if err != nil {
		log.Fatal(err)
	}

	for _, post := range posts.GetPosts() {
		fmt.Println(post.Id)
		fmt.Println(post.Title)
		fmt.Println(post.Text)
	}

}
