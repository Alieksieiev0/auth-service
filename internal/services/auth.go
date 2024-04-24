package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Alieksieiev0/auth-service/api/proto"
	"github.com/Alieksieiev0/auth-service/internal/types"
	"golang.org/x/crypto/bcrypt"
)

const exp time.Duration = time.Hour * 6

type AuthService interface {
	Register(ctx context.Context, user *types.User) error
	Login(ctx context.Context, user *types.User) (*types.UserToken, error)
	ReadClaims(ctx context.Context, token *types.Token) (*types.UserClaims, error)
}

type authService struct {
	client       proto.UserServiceClient
	tokenService TokenService
}

func NewAuthService(client proto.UserServiceClient, tokenService TokenService) AuthService {
	return &authService{
		client:       client,
		tokenService: tokenService,
	}
}

func (as *authService) Register(ctx context.Context, user *types.User) error {
	userRequest := &proto.UserRequest{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	user.Password = ""
	_, err := as.client.Save(ctx, userRequest)
	return err
}

func (as *authService) Login(ctx context.Context, user *types.User) (*types.UserToken, error) {
	username := &proto.UsernameRequest{
		Username: user.Username,
	}

	resp, err := as.client.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("user with such username does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf("provided password is incorrect")
	}

	respUser := &types.User{
		Id:       resp.Id,
		Username: resp.Username,
		Email:    resp.Email,
	}
	token, err := as.tokenService.Create(respUser, exp)
	if err != nil {
		return nil, err
	}

	userToken := &types.UserToken{
		User: *respUser,
		Token: types.Token{
			Value: token,
		},
	}

	return userToken, nil
}

func (as *authService) ReadClaims(
	ctx context.Context,
	token *types.Token,
) (*types.UserClaims, error) {
	fmt.Println(as.tokenService)
	claims, err := as.tokenService.Read(token.Value)
	if err != nil {
		return nil, err
	}

	userClaims, ok := claims.(*types.UserClaims)
	if !ok {
		return nil, fmt.Errorf("you're Unauthorized due to claims parsing error")
	}

	return userClaims, nil
}
