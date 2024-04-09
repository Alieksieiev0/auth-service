package grpc

import (
	"context"
	"net"

	"github.com/Alieksieiev0/auth-service/api/proto"
	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/Alieksieiev0/auth-service/internal/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCServer struct {
}

func NewServer() *GRPCServer {
	return &GRPCServer{}
}

func (as *GRPCServer) Start(addr string, service services.AuthService) error {
	grpcUserService := NewGRPCAuthServiceServer(service)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	proto.RegisterUserServiceServer(server, grpcUserService)

	return server.Serve(ln)
}

type GRPCAuthServiceServer struct {
	service services.AuthService
	proto.UnimplementedUserServiceServer
}

func NewGRPCAuthServiceServer(service services.AuthService) *GRPCAuthServiceServer {
	return &GRPCAuthServiceServer{
		service: service,
	}
}

func (as *GRPCAuthServiceServer) Register(
	ctx context.Context,
	req *proto.User,
) (*emptypb.Empty, error) {
	user := &types.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	err := as.service.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

func (as *GRPCAuthServiceServer) Login(
	ctx context.Context,
	req *proto.User,
) (*proto.Token, error) {
	user := &types.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	token, err := as.service.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return &proto.Token{Value: token.Value}, err
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
		ExpiresAt: timestamppb.New(userClaims.ExpiresAt.Time),
		NotBefore: timestamppb.New(userClaims.NotBefore.Time),
		IssuedAt:  timestamppb.New(userClaims.IssuedAt.Time),
		Id:        userClaims.ID,
	}
	return claims, err
}
