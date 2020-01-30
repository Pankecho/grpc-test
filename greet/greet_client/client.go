package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"github.com/Pankecho/grpc-test/greet/greetpb"
	"log"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect due to: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	request := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{ FirstName: "Juan Pablo", LastName: "Martínez Ruiz" },
	}

	resp, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Couldnt make the request: %v", err)
	}
	fmt.Println(resp.Result)
}
