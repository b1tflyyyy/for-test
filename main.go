package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grpcSRV := grpc.NewServer()

	srv := NewNumbersService()
	srv = NewLoggerService(srv)

	RegisterNumbersStorageServer(grpcSRV, srv)
	listener, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatal(err)
	}

	if err := grpcSRV.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
