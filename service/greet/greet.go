package greet

import (
	"context"
	"fmt"

	"github.com/ZAF07/go-microservice-prac/constants"
	"github.com/ZAF07/go-microservice-prac/manager"
	"github.com/ZAF07/go-microservice-prac/proto/rpc"
)

// Define a service struct to represent the service
// Service struct has to implement all methods defined in the GRPC interface
type Greet struct {
	rpc.UnimplementedGreetServiceServer
}

// This gurantees that Service struct implements the interface
var _ rpc.GreetServiceServer = (*Greet)(nil)

// Method to initialise a new service
// Returns a pointer to the service struct
func NewGreetService() *Greet {
	return&Greet{}
}

// Method from GRPC interface
func (s *Greet) Greetings(ctx context.Context, req *rpc.Greeting) (resp *rpc.Reply, err error) {
	fmt.Println("RECIEVED A REQUEST !!")
	fmt.Printf("CONTEXT RECEIVED : %+v\n", ctx.Value(constants.CtxKey("name")))
	fmt.Println("SAVING FILE...")
	saveErr := manager.Save()
	if saveErr != nil {
		fmt.Println("SAVE ERROR")
	} else {
		fmt.Println("Save success")
	}
// this
	manager.Shout()
	resp = &rpc.Reply{
		Reply: "Yes hello",
	}
	return
}