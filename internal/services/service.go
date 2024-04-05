package services

import (
	"context"
	"fmt"
)

type AuthService interface {
	Register(ctx context.Context) error
	Login(ctx context.Context) error
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (as *authService) Login(ctx context.Context) error {
	fmt.Println(1)
	return nil
}

func (as *authService) Register(ctx context.Context) error {
	fmt.Println(2)
	return nil
}
