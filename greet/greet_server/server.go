package main

import (
	"context"
	"google.golang.org/grpc"
	"github.com/Pankecho/grpc-test/greet/greetpb"
	"log"
	"net"
)

type Server struct {}

func (*Server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	response := "Hello " + " " + firstName
	return &greetpb.GreetingResponse{
		Result: response,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
