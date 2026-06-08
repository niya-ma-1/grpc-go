package main

import (
	"context"
	"log"
	"time"

	pb "github.com/niya-ma-1/grpc-go/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked.")

	reqs := []float64{1,2,3,4}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("error while calling Average: %v", err)
	}

	for _, number := range reqs {
		log.Printf("Sending req: %v\n", number)
		stream.Send(&pb.AverageRequest{
			Number: number,
		})
		
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from Average: %v", err)
	}

	log.Printf("Average response: %v", res.Result)
}
