package main

import (
	"log"
	"net"

	"github.com/elewis787/streamingrpc/pb"
	"github.com/elewis787/streamingrpc/server"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterTransportServer(grpcServer, &server.Server{})
	lis, err := net.Listen("tcp", "127.0.0.1:8787")
	if err != nil {
		log.Println(err)
	}
	log.Println("Starting server...")
	// log running
	if err := grpcServer.Serve(lis); err != nil {
		log.Println(err)
	}
}
