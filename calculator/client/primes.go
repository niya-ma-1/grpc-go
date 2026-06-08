package main

import (
	"context"
	"log"
	"io"

	pb "github.com/niya-ma-1/grpc-go/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: 1239094,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes RPC: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving the stream: %v\n", err)
		}

		log.Printf("Prime: %d\n", res.Result)
	}
}
