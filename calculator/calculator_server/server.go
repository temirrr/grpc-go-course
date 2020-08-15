package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/temirrr/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	k := int64(2)
	N := req.GetNumber()
	for N > 1 {
		if N%k == 0 {
			N = N / k
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Number: k,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
		} else {
			k++
		}
	}
	return nil
}

func main() {
	fmt.Println("Server has started.")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
