package main

import (
	"log"
	"net"

	pb "github.com/klimenko-serj/grpc-test/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	log.Println("UrlServiceServer started")

	l, err := net.Listen("tcp", ":9099")
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
