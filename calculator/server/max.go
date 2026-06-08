package main

import (
	"io"
	"log"

	pb "github.com/niya-ma-1/grpc-go/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked.")
	max_result := int32(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if req.Number > max_result {
			max_result = max(max_result, req.Number)
			err = stream.Send(&pb.MaxResponse{
				Result: max_result,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
			}
		}

	}
}
