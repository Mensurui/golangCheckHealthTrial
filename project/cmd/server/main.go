package main

import (
	"net"

	"github.com/Mensurui/golangCheckHealthTrial/project/internal"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	protos "github.com/Mensurui/golangCheckHealthTrial/protos/golang"
)

func main() {
	logger := hclog.Default()
	healthService := internal.NewService(logger)

	grpcServer := grpc.NewServer()
	protos.RegisterServiceServer(grpcServer, healthService)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":9092")
	if err != nil {
		logger.Warn("[ERROR LISTENING ON PORT :9092]")
	}

	logger.Info("Starting gRPC server on port 9092")

	if err := grpcServer.Serve(listener); err != nil {
		logger.Error("Failed to server", "error", err)
	}
}
