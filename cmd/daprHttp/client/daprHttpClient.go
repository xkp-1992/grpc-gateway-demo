package main

import (
	"context"
	dapr "github.com/dapr/go-sdk/client"
	"log"
	"time"
)

const (
	address = "localhost:8500"
)

func main() {
	ctx := context.Background()

	// create the dapr
	client, err := dapr.NewClientWithAddressContext(context.Background(), address)
	//client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	for {
		resp, err := client.InvokeMethod(ctx, "http-server", "hello", "get")
		if err != nil {
			panic(err)
		}
		log.Printf("go-service method hello has invoked, response is: %s", string(resp))
		time.Sleep(time.Second * 5)
	}
}
