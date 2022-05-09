package main

import (
	"fmt"
	"net"

	tickets "github.com/postie-labs/go-tickets/app"
	"github.com/postie-labs/go-tickets/server"
	pb "github.com/postie-labs/proto/billetterie"
	"google.golang.org/grpc"
)

const (
	DefaultHost = "localhost"
	DefaultPort = "7788"
)

func main() {
	fmt.Println("things happen from here")

	listener, err := net.Listen("tcp", net.JoinHostPort(DefaultHost, DefaultPort))
	if err != nil {
		panic(err)
	}

	// init app
	app, err := tickets.NewApplication()
	if err != nil {
		panic(err)
	}

	// init server(handler)
	billetterieServer := server.NewBilletterieServer(app)

	// init grpcServer
	opts := []grpc.ServerOption{} // TODO: add some options
	grpcServer := grpc.NewServer(opts...)

	// register server(handler) to grpcServer
	pb.RegisterBilletterieServer(grpcServer, billetterieServer)

	// start grpcServer
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
