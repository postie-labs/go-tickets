package server

import (
	"context"

	"github.com/postie-labs/go-postie-lib/crypto"
	tickets "github.com/postie-labs/go-tickets/app"
	"github.com/postie-labs/go-tickets/types"
	pb "github.com/postie-labs/proto/billetterie"
)

type BilletterieServer struct {
	pb.UnimplementedBilletterieServer

	app *tickets.Application
}

func NewBilletterieServer(app *tickets.Application) *BilletterieServer {
	return &BilletterieServer{app: app}
}

func (s *BilletterieServer) Issue(ctx context.Context, req *pb.IssueRequest) (*pb.IssueResponse, error) {
	// prepare
	issuerAddr := crypto.Addr(req.Issuer)
	data := req.Data

	// TODO: check

	// execute
	tckHash, err := s.app.Issue(issuerAddr, data)
	if err != nil {
		return nil, err
	}

	return &pb.IssueResponse{TicketHash: tckHash.String()}, nil
}

func (s *BilletterieServer) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {
	// prepare
	ownerFrom := crypto.Addr(req.OwnerFrom)
	ownerTo := crypto.Addr(req.OwnerTo)
	tckHash, err := types.NewHashFromString(req.TicketHash)
	if err != nil {
		return nil, err
	}

	// TODO: check

	// execute
	err = s.app.Transfer(ownerFrom, ownerTo, tckHash)
	if err != nil {
		return nil, err
	}

	return &pb.TransferResponse{Transferred: true}, nil
}

func (s *BilletterieServer) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	// prepare
	ownerAddr := crypto.Addr(req.Owner)
	tckHash, err := types.NewHashFromString(req.TicketHash)
	if err != nil {
		return nil, err
	}

	// TODO: check

	// execute
	verified, err := s.app.Verify(ownerAddr, tckHash)
	if err != nil {
		return nil, err
	}

	return &pb.VerifyResponse{Verified: verified}, nil
}
