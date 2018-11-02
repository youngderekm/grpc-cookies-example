// gRPC service

package main

import (
	"fmt"
	"net"
	"os"

	"github.com/youngderekm/grpc-cookies-example/servicedef"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// service is the struct that our gRPC service methods will be bound to
type server struct {
}

func main() {

	// Create the grpc server
	grpcServer := grpc.NewServer()

	authServer := server{}

	// Register our service
	servicedef.RegisterAuthApiServer(grpcServer, &authServer)

	// Set up the listener
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Printf("API server failed on net.Listen: %v\n", err)
		os.Exit(1)
	}

	reflection.Register(grpcServer) // Register reflection service on gRPC api.
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("API server failed on grpc Serve: %v\n", err)
		os.Exit(1)
	}
}
