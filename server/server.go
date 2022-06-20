package main

import (
	"context"
	"log"
	"net"

	pb "github.com/tboerc/gwall/protos"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeGreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}

// Greet greets with FirstName
func (*server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	result := "Hello " + in.GetGreeting().GetFirstName()
	res := pb.GreetResponse{
		Result: result,
	}
	return &res, nil
}
