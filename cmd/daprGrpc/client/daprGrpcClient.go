package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pb "www.github.com/grpc-gateway-demo/gatewayProtos"
)

const (
	address = "localhost:5501"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "grpc-server")
	for {
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Dapr", Message: "Dapr Message"})
		if err != nil {
			log.Printf("could not hello world: %v", err)
			log.Println()
		} else {
			log.Printf("hello world: %s	%s", r.Name, r.Message)
			log.Println()
		}
		time.Sleep(time.Second * 1)
	}
}
