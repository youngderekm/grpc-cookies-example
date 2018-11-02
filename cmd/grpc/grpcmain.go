// gRPC service

package main

import (
	"net"
	"os"

	"github.com/romana/rlog"
	"github.com/youngderekm/grpc-cookies-example/servicedef"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// service is the struct that our gRPC service methods will be bound to
type server struct {
}

func main() {
	// log at debug for this demo
	os.Setenv("RLOG_LOG_LEVEL", "DEBUG")
	rlog.UpdateEnv()

	// Create the grpc server
	grpcServer := grpc.NewServer()

	authServer := server{}

	// Register our service
	servicedef.RegisterAuthApiServer(grpcServer, &authServer)

	// Set up the listener
	hostAndPort := "localhost:50051"
	lis, err := net.Listen("tcp", hostAndPort)
	if err != nil {
		rlog.Criticalf("API server failed on net.Listen: %v", err)
		os.Exit(1)
	}

	rlog.Infof("Listening for gRPC requests at %s", hostAndPort)

	reflection.Register(grpcServer) // Register reflection service on gRPC api.
	if err := grpcServer.Serve(lis); err != nil {
		rlog.Criticalf("API server failed on grpc Serve: %v", err)
		os.Exit(1)
	}
}
