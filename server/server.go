package server

import (
	"context"

	pb "github.com/postie-labs/go-postie-proto/tickets"
)

type BilletterieServer struct {
	pb.UnimplementedBilletterieServer
}

func NewBilletterieServer() *BilletterieServer {
	return &BilletterieServer{}
}

func (s *BilletterieServer) Issue(context.Context, *pb.IssueRequest) (*pb.IssueResponse, error) {
	return nil, nil
}

func (s *BilletterieServer) Transfer(context.Context, *pb.TransferRequest) (*pb.TransferResponse, error) {
	return nil, nil
}

func (s *BilletterieServer) Verify(context.Context, *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	return nil, nil
}
