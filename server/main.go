package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/tboerc/gwall/protos"
	"google.golang.org/grpc"
)

func Getenv(key string, fallback string) string {
	env, exist := os.LookupEnv(key)
	if exist {
		return env
	}
	return fallback
}

func main() {
	PORT := Getenv("PORT", "8080")
	HOST := Getenv("HOST", "localhost")

	addr := fmt.Sprintf("%s:%s", HOST, PORT)

	log.Println("starting server on", addr)

	s := grpc.NewServer()
	pb.RegisterRoomServer(s, &server{})

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error: failed to listen %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error: failed to start server %v", err)
	}
}
