package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/postie-labs/go-postie-lib/crypto"
	pb "github.com/postie-labs/proto/tickets"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DefaultHost = "localhost"
	DefaultPort = "7788"
)

func main() {
	// init grpcDial, TODO: add some options
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(net.JoinHostPort(DefaultHost, DefaultPort), opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// init client
	client := pb.NewBilletterieClient(conn)

	// init private keys
	alice, err := crypto.GenPrivKey()
	if err != nil {
		panic(err)
	}
	aliceAddr := alice.PubKey().Address()
	bob, err := crypto.GenPrivKey()
	if err != nil {
		panic(err)
	}
	bobAddr := bob.PubKey().Address()

	// init context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// init data
	data := []byte("hello world")

	// (alice) ISSUE (ticket A)
	issueRes, err := client.Issue(ctx, &pb.IssueRequest{
		Issuer: string(aliceAddr),
		Data:   data,
	})
	if err != nil {
		panic(err)
	}
	tckHash := issueRes.TicketHash
	fmt.Printf("issued %s\n", tckHash)

	// (alice) TRANNSFER (ticket A) to (bob)
	transferRes, err := client.Transfer(ctx, &pb.TransferRequest{
		OwnerFrom:  string(aliceAddr),
		OwnerTo:    string(bobAddr),
		TicketHash: tckHash,
	})
	if err != nil {
		panic(err)
	}
	if !transferRes.Transferred {
		panic("failed to transfer ownership")
	}
	fmt.Printf("transferred %s from %s to %s\n", tckHash, aliceAddr, bobAddr)

	// (alice) VERIFY (ticket A): false
	verifyRes, err := client.Verify(ctx, &pb.VerifyRequest{
		Owner:      string(aliceAddr),
		TicketHash: tckHash,
	})
	if err != nil {
		panic(err)
	}
	if verifyRes.Verified {
		panic("failed to verify ownership")
	}
	fmt.Printf("verified %s NOT to own %s\n", aliceAddr, tckHash)

	// (bob) VERIFY (ticket A): true
	verifyRes, err = client.Verify(ctx, &pb.VerifyRequest{
		Owner:      string(bobAddr),
		TicketHash: tckHash,
	})
	if err != nil {
		panic(err)
	}
	if !verifyRes.Verified {
		panic("failed to verify ownership")
	}
	fmt.Printf("verified %s to own %s\n", bobAddr, tckHash)

	// (bob) TRANSFER (ticket A) to (alice)
	transferRes, err = client.Transfer(ctx, &pb.TransferRequest{
		OwnerFrom:  string(bobAddr),
		OwnerTo:    string(aliceAddr),
		TicketHash: tckHash,
	})
	if err != nil {
		panic(err)
	}
	if !transferRes.Transferred {
		panic("failed to transfer ownership")
	}
	fmt.Printf("transferred %s from %s to %s\n", tckHash, bobAddr, aliceAddr)

	// (alice) VERIFY (ticket A): true
	verifyRes, err = client.Verify(ctx, &pb.VerifyRequest{
		Owner:      string(aliceAddr),
		TicketHash: tckHash,
	})
	if err != nil {
		panic(err)
	}
	if !verifyRes.Verified {
		panic("failed to verify ownership")
	}
	fmt.Printf("verified %s to own %s\n", aliceAddr, tckHash)

	// (bob) VERIFY (ticket A): false
	verifyRes, err = client.Verify(ctx, &pb.VerifyRequest{
		Owner:      string(bobAddr),
		TicketHash: tckHash,
	})
	if err != nil {
		panic(err)
	}
	if verifyRes.Verified {
		panic("failed to verify ownership")
	}
	fmt.Printf("verified %s NOT to own %s\n", bobAddr, tckHash)
}
