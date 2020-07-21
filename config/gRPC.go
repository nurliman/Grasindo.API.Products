package config

import (
	"fmt"
	"log"
	"net"

	"github.com/nurliman/Grasindo.API.Products/protos"
	"google.golang.org/grpc"
)

func startGRPC() {

	fmt.Println("Products Service starting")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("on port", 9000)

	s := protos.Server{}

	grpcServer := grpc.NewServer()

	protos.RegisterProductsServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
