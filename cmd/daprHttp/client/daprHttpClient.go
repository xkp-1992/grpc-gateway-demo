package main

import (
	"context"
	"fmt"
	"github.com/dapr/go-sdk/client"
	"log"
	"time"
)

const (
	address = "localhost:8500"
)

func main() {
	ctx := context.Background()
	//create the dapr
	c, err := client.NewClientWithAddressContext(ctx, address) //无法调用到接口
	//c, err := client.NewClient()				//调用成功
	if err != nil {
		panic(err)
	}
	defer c.Close()
	for {
		resp, err := c.InvokeMethod(ctx, "http-server", "hello", "get")
		if err != nil {
			fmt.Println("error start")
			fmt.Println(err.Error())
			fmt.Println("error end")
			//panic(err)
		} else {
			log.Printf("go-service method hello has invoked, response is: %s", string(resp))
		}
		//log.Printf("go-service method hello has invoked, response is: %s", string(resp))
		time.Sleep(time.Second * 5)
	}
}
