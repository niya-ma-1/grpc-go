package main

import (
	"context"
	"io"
	"log"

	pb "github.com/niya-ma-1/grpc-go/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Niya",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", res.Result)
	}

}
