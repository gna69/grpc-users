package main

import (
	"github.com/gna69/grpc-users/internal/server"
	"github.com/gna69/grpc-users/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	srv := server.New()

	api.RegisterUserServer(grpcServer, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
