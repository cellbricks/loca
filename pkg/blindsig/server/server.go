package server

import (
	"context"
	"sync"

	pb "cellbricks/LoA/pkg/blindsig/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type senderServer struct {
	pb.SenderServer

	mu sync.Mutex
}

func NewSenderServer() *senderServer {
	s := &senderServer{}
	return s
}

func (s *senderServer) Gen(context.Context, *pb.GenRequest) (*pb.GenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gen not implemented")
}

type signerServer struct {
	pb.SignerServer

	mu sync.Mutex
}

func NewSignerServer() *signerServer {
	s := &signerServer{}
	return s
}

func (s *signerServer) Sign(context.Context, *pb.SignRequest) (*pb.SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}

func (s *signerServer) Verify(context.Context, *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
