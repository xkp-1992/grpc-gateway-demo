package main

import (
	"log"
	"net"

	daprd "github.com/dapr/go-sdk/service/grpc"
	"google.golang.org/grpc"
	"www.github.com/grpc-gateway-demo/gatewayProtos"
	"www.github.com/grpc-gateway-demo/grpcServers"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	gatewayProtos.RegisterHelloServiceServer(s, &grpcServers.XHelloWorld{})
	daprServer := daprd.NewServiceWithGrpcServer(lis, s)

	// start the server
	if err := daprServer.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
