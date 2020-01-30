package main

import (
	"context"
	"google.golang.org/grpc"
	"github.com/Pankecho/grpc-test/calculator/sumpb"
	"log"
	"net"
)

type CalculatorServer struct {}

func (*CalculatorServer) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	x := req.GetSum().GetX()
	y := req.GetSum().GetY()
	response := sumpb.SumResponse{Sum: x + y }
	return &response, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &CalculatorServer{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
