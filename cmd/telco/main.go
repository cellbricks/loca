package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "cellbricks/LoA/pkg/blindsig/proto"
	blindsig "cellbricks/LoA/pkg/blindsig/server"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSenderServer(grpcServer, blindsig.NewSenderServer())
	grpcServer.Serve(lis)
}

