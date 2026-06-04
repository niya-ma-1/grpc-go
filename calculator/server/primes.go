package main

import (
	"log"

	pb "github.com/niya-ma-1/grpc-practice/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	divisor := int64(2)
	number := in.Number

	for number > 1 {
		if number%divisor == 0 { // if divisor evenly divides into number
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number /= divisor // divide number by divisor so that we have the rest of the number left.
		} else {
			divisor = divisor + 1
		}
	}

	return nil
}
