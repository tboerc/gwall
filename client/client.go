package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/tboerc/gwall/shared"
	pb "github.com/tboerc/gwall/shared/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
	url := shared.Getenv("API_URL", "http://localhost:8080")

	body, _ := proto.Marshal(&pb.CreateUserRequest{Username: "tboerc", Password: []byte("Teste")})

	resp, _ := http.Post(url+"/v1/user/create", "application/protobuf", bytes.NewBuffer(body))

	log.Println(resp.StatusCode)
}
