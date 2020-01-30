package main

import (
	"context"
	"fmt"
	"github.com/Pankecho/grpc-test/calculator/sumpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect due to: %v", err)
	}
	defer conn.Close()

	c := sumpb.NewSumServiceClient(conn)

	request := &sumpb.SumRequest{
		Sum: &sumpb.Sum{X: 5, Y: 10},
	}

	resp, err := c.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Couldnt make the request: %v", err)
	}
	fmt.Println(resp.Sum)
}
