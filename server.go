package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/klimenko-serj/grpc-test/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/grpc/grpclog"
)

func main() {
	log.Println("UrlServiceServer started")

	serverPort := os.Getenv("GRPC_TEST_SERVER_PORT")
	if serverPort == "" {
		serverPort = "9099"
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUrlServiceServer(grpcServer, &server{})

	grpcServer.Serve(l)
}


type server struct {}

func (s* server) ProcessURL(ctx context.Context,u *pb.UrlRequest) (*emptypb.Empty, error){
	log.Println("UrlReqest:", u.Url)

	p, _ := peer.FromContext(ctx)
	clientIP := p.Addr.(*net.TCPAddr).IP.String()
	log.Println("Client IP:", clientIP)

	go processURL(clientIP, u.Url)

	return &emptypb.Empty{}, nil
}

func processURL(clientIP, url string) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	clientPort := os.Getenv("GRPC_TEST_CLIENT_PORT")
	if clientPort == "" {
		clientPort = "9077"
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", clientIP, clientPort), opts...)
	if err != nil {
		// No Fatal - it shouldn't stop server
		grpclog.Errorf("GRPC Dial failed: %v", err)
		return
	}

	defer conn.Close()

	client := pb.NewUrlClientClient(conn)

	

	request := &pb.Header{StatusCode: 200}

	_, err = client.SendHeader(context.Background(), request)
	if err != nil {
		// No Fatal - it shouldn't stop server
		grpclog.Errorf("GRPC call ProcessURL failed: %v", err)
		return
	}
	log.Println("Header sent")

	_, err = client.Finish(context.Background(), &emptypb.Empty{})
	if err != nil {
		// No Fatal - it shouldn't stop server
		grpclog.Errorf("GRPC call Finish failed: %v", err)
		return
	}
	log.Println("Finish signal sent")
}