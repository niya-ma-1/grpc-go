package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/niya-ma-1/grpc-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked.")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("error while creating stream: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{1, 5, 3, 6, 2, 20}
		for _, number := range numbers {
			log.Printf("Sending req: %v\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
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
