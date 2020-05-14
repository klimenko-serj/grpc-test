package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/klimenko-serj/grpc-test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	file, err := os.OpenFile("client.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)

	if len(os.Args) < 2 {
		fmt.Println("Usage:\n\tclient [url]")
		return
	}
	url := os.Args[1]

	serverAddr := os.Getenv("GRPC_TEST_SERVER")
	if serverAddr == "" {
		serverAddr = "127.0.0.1:9099"
	}

	finished := make(chan bool)
	go startClientService(finished)

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

	<-finished
}

func startClientService(f chan bool) {
	serverPort := os.Getenv("GRPC_TEST_CLIENT_PORT")
	if serverPort == "" {
		serverPort = "9077"
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	clientServ := client{grpcServer: grpcServer, bodyBuffer: &bytes.Buffer{}}

	pb.RegisterUrlClientServer(grpcServer, &clientServ)

	grpcServer.Serve(l)
	log.Println("gRPC Server stopped.")

	strBody := string(clientServ.bodyBuffer.Bytes())
	strBodyLen := len(strBody)

	log.Println("BodyBuffer size: ", strBodyLen)

	if strBodyLen > 100 {
		fmt.Println(strBody[:100])
	} else {
		fmt.Println(strBody)
	}
	f <- true
}

type client struct {
	grpcServer *grpc.Server
	bodyBuffer *bytes.Buffer
}

func (c client) SendHeader(ctx context.Context, h *pb.Header) (*empty.Empty, error) {
	log.Println("StatusCode:", h.StatusCode)
	log.Println("Header:\n", h.Header)
	fmt.Println(h.Header)
	return &empty.Empty{}, nil
}

func (c client) SendBody(ctx context.Context, b *pb.Body) (*empty.Empty, error) {
	log.Println("Body part: size =", len(b.Body))
	c.bodyBuffer.Write(b.Body)
	return &empty.Empty{}, nil
}

func (c *client) Finish(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	go func() {
		c.grpcServer.GracefulStop()
	}()
	return &empty.Empty{}, nil
}
