package main

import (
	"log"
	"context"
	"time"

	pb "github.com/niya-ma-1/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked.")

	reqs := []*pb.GreetRequest{
		{FirstName: "Niya"},
		{FirstName: "Mia"},
		{FirstName: "Test"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet response: %v", res.Result)
}