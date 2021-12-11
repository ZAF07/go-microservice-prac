package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ZAF07/go-microservice-prac/proto/rpc"
	"google.golang.org/grpc"
)

// This mimics a client cal from another service
// Wil implement diff ports later

func main() { 
	// Create a connection instance and Dial the GRPC server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to GRPC client : %+v", err)
	}
	defer conn.Close()

	// Initialises a new GRPC client service stub
	c := rpc.NewGreetServiceClient(conn)

	// Construct the request body to pass in GRPC Service method
	req := &rpc.Greeting{
	Name: "Zaffere",
	}
	ctx := context.Background()

	// Invoke the GRPC Service method and wait for response (Unary)
	resp, rErr := c.Greetings(ctx, req)
	if err != nil {
		log.Fatalf("bad response : %+v", rErr)
	}

	fmt.Println(resp)
	
}