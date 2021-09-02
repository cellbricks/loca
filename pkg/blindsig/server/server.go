package server

import (
	"flag"
	"sync"

	pb "cellbricks/LoA/pkg/blindsig/proto"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type senderServer struct {
	pb.SenderServer

	mu sync.Mutex // protects routeNotes
}

func NewSenderServer() *senderServer {
	s := &senderServer{}
	return s
}

