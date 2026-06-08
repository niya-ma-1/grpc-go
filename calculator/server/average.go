package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/niya-ma-1/grpc-go/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Print("Average function was invoked.")

	res := float64(0)
	count := float64(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(res / count),
			})
		}

		if err != nil {
			return fmt.Errorf("error while reading client stream: %w", err)
		}

		log.Printf("Receiving: %v\n", req.Number)

		res += float64(req.Number)
		count += float64(1)
	}
}
