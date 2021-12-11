package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	rpc "github.com/ZAF07/go-microservice-prac/proto/rpc"
	"github.com/ZAF07/go-microservice-prac/router"
	"github.com/ZAF07/go-microservice-prac/service/greet"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)


func main() {
	// Spin up the main server instance
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("cannot connect to tcp port 8000 : %+v", err)
	}

	// Start a new multiplexer passing in the main server
	m := cmux.New(lis)

	// Listen for HTTP requests first
	// If request headers don't specify HTTP, next mux would handle the request
	httpListener := m.Match(cmux.HTTP1Fast())
	grpclistener := m.Match(cmux.Any())
	
	// Run GO routine to run both processes at the same time
	go serveGRPC(grpclistener)
	go serveHTTP(httpListener)

	fmt.Printf("Running@%v\n", lis.Addr())

	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		log.Fatalf("MUX ERR : %+v", err)
	}

}

// GRPC Server initialisation
func serveGRPC(l net.Listener) {
	grpcServer := grpc.NewServer()
	
	rpc.RegisterGreetServiceServer(grpcServer, greet.NewGreetService())
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("error running GRPC server %+v", err)
	}
}

// HTTP Server initialisation (using gin)
func serveHTTP(l net.Listener) {
	h := gin.Default()
	router.Router(h)
	s := &http.Server{
		Handler: h,
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		log.Fatalf("error serving HTTP : %+v", err)
	}
}