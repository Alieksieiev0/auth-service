package grpc

import (
	"github.com/Alieksieiev0/auth-service/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewConnection(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewGRPCUserServiceClient(addr string) (proto.UserServiceClient, error) {
	conn, err := NewConnection(addr)
	if err != nil {
		return nil, err
	}

	return proto.NewUserServiceClient(conn), nil
}
