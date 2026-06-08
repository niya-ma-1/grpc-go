package main

import (
	"context"
	"log"
	"time"

	pb "github.com/niya-ma-1/grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Niya",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
			} else {
				log.Fatalf("unexpected gRPC error: %v\n", e)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
		return
	}

	log.Printf("GreetWithDeadline response: %s\n", res.Result)

}
