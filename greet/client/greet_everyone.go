package main

import (
	"context"
	"log"
	"time"
	"io"

	pb "github.com/niya-ma-1/grpc-go/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked.")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("error while creating stream: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Niya"},
		{FirstName: "Mia"},
		{FirstName: "Test"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("error while receiving response: %v\n", err)
			}

			log.Printf("Received: %v\n", res)
		}
		close(waitc)
	}()

	<-waitc

}