package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	pb "cellbricks/LoA/pkg/blindsig/proto"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type blindSigServer struct {
	pb.UnimplementedBlindSigServer

	mu sync.Mutex // protects routeNotes
}

func newServer() *blindSigServer {
	s := &blindSigServer{}
	return s
}

func (s *blindSigServer) GetToken(ctx context.Context, msg *pb.Message) (*pb.Token, error) {
	// TBD
	//t := *pb.Token{
	//	token: "",
	//}

	return nil, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBlindSigServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
