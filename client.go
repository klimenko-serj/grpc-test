package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/klimenko-serj/grpc-test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n\tclient [url]")
		return
	}
	url := os.Args[1]

	serverAddr := os.Getenv("GRPC_TEST_SERVER")
	if serverAddr == "" {
		serverAddr = "127.0.0.1:9099"
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(serverAddr, opts...)

	if err != nil {
		grpclog.Fatalf("GRPC Dial failed: %v", err)
	}

	defer conn.Close()

	client := pb.NewUrlServiceClient(conn)
	request := &pb.UrlRequest{
		Url: url,
	}

	_, err = client.ProcessURL(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("GRPC call ProcessURL failed: %v", err)
	}
	log.Println("Request to process", url, "sent.")

}
