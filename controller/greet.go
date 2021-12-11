package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ZAF07/go-microservice-prac/constants"
	"github.com/ZAF07/go-microservice-prac/proto/rpc"
	"github.com/ZAF07/go-microservice-prac/service/greet"
	"github.com/gin-gonic/gin"
)

// Define struct to repersent the controller
type GreetAPI struct {}

func (api GreetAPI) SayHello(c *gin.Context) {
	fmt.Println("GETTING THERE ZAFFERE !!")
	name := "Zaffere"
	// Create a context to pass to service layer
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.CtxKey("name"), name)


	// Initialise a new service instance
	service := greet.NewGreetService()

	// Prepare the req to send to GRPC service
	req := &rpc.Greeting{
		Name: "Jolly Joy",
	}

	// Invoke the service method and wait for response (Unary)
	resp, err := service.Greetings(ctx, req)
	if err != nil {
		log.Fatalf("error invoking the service : %+v", err)
	}
	
	// Send response back to client
	c.JSON(http.StatusOK,
	gin.H{
		"message": "Success",
		"status": http.StatusOK,
		"data": resp,
	})
}