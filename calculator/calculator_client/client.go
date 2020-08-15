package main

import (
	"fmt"
	"io"
	"log"

	"github.com/temirrr/grpc-go-course/calculator/calculatorpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client program has been started.")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	doServerStreaming(c)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 138, // Just some random number, which could have been taken as argument in CLI.
	}
	conn, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not invoke the Prime Number Decomposition Service: %v", err)
	}

	for {
		res, err := conn.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("One of the divisors is: %d\n", res.GetNumber())
	}
}
