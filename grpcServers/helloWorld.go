package grpcServers

import (
	"context"
	"fmt"
	"www.github.com/grpc-gateway-demo/gatewayProtos"
)

type XHelloWorld struct {
	gatewayProtos.UnimplementedHelloServiceServer
}

func (xHelloWorld *XHelloWorld) SayHello(ctx context.Context, in *gatewayProtos.HelloRequest) (*gatewayProtos.HelloResponse, error) {

	fmt.Println(in.Name)
	fmt.Println(in.Message)
	return &gatewayProtos.HelloResponse{Message: fmt.Sprintf(in.Message, "response message"), Name: fmt.Sprintf(in.Name, "response name")}, nil
}
