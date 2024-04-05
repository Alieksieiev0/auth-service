package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Alieksieiev0/auth-service/api/proto"
	"github.com/Alieksieiev0/auth-service/internal/transport/grpc"
)

func main() {
	fmt.Println(1)
	client, err := grpc.NewGRPCClient(":4000")
	fmt.Println(2)
	if err != nil {
		log.Fatal(err)
	}

	usr := proto.UsernameRequest{
		Username: "test",
	}
	user, err := client.GetByUsername(context.Background(), &usr)
	fmt.Println(user)
	if err != nil {
		log.Fatal(err)
	}
}
