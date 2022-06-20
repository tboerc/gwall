package main

import (
	"context"
	"log"

	pb "github.com/tboerc/gwall/protos"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pb.UnimplementedRoomServer
}

func (*server) Join(ctx context.Context, in *pb.JoinRequest) (*pb.JoinResponse, error) {
	uuid := in.GetUuid()
	md, _ := metadata.FromIncomingContext(ctx)

	ip := md.Get(":authority")[0]

	log.Println("Ip is", ip)
	log.Println("UUID is", uuid)

	res := pb.JoinResponse{}

	return &res, nil
}
