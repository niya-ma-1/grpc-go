package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/niya-ma-1/grpc-go/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Print("LongGreet function was invoked.")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			return fmt.Errorf("error while reading client stream: %w", err)
		}

		log.Printf("Receiving: %v\n", req)

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
