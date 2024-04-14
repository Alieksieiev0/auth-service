package grpc

import (
	"context"
	"net"

	"github.com/Alieksieiev0/auth-service/api/proto"
	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/Alieksieiev0/auth-service/internal/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCServer struct {
	addr string
}

func NewServer(addr string) *GRPCServer {
	return &GRPCServer{
		addr: addr,
	}
}

func (s *GRPCServer) Start(service services.AuthService) error {
	grpcAuthServiceServer := NewGRPCAuthServiceServer(service)

	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	proto.RegisterAuthServiceServer(server, grpcAuthServiceServer)

	return server.Serve(ln)
}

type GRPCAuthServiceServer struct {
	service services.AuthService
	proto.UnimplementedAuthServiceServer
}

func NewGRPCAuthServiceServer(service services.AuthService) *GRPCAuthServiceServer {
	return &GRPCAuthServiceServer{
		service: service,
	}
}

func (as *GRPCAuthServiceServer) ReadClaims(
	ctx context.Context,
	req *proto.Token,
) (*proto.UserClaims, error) {
	token := &types.Token{
		Value: req.Value,
	}
	userClaims, err := as.service.ReadClaims(ctx, token)
	if err != nil {
		return nil, err
	}

	claims := &proto.UserClaims{
		UserId:    userClaims.UserId,
		Username:  userClaims.Username,
		Email:     userClaims.Email,
		Algorithm: userClaims.Algorithm,
		Issuer:    userClaims.Issuer,
		Subject:   userClaims.Subject,
		Audience:  userClaims.Audience,
		Id:        userClaims.ID,
	}

	if userClaims.ExpiresAt != nil {
		claims.ExpiresAt = timestamppb.New(userClaims.ExpiresAt.Time)
	}

	if userClaims.NotBefore != nil {
		claims.NotBefore = timestamppb.New(userClaims.NotBefore.Time)
	}

	if userClaims.IssuedAt != nil {
		claims.IssuedAt = timestamppb.New(userClaims.IssuedAt.Time)
	}

	return claims, err
}
