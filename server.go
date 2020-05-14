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
	addr := p.Addr.(*net.TCPAddr).IP.String()
	log.Println("Addr:", addr)

	return &emptypb.Empty{}, nil
}
