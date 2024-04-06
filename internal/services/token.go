package services

import (
	"fmt"
	"time"

	"github.com/Alieksieiev0/auth-service/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

const secret string = "d0881af56a560c2b86c7dd95ee5bd95a52864e2bd0b42396bba0b5ce974fe2cb6065eb3b03776997a6e9c2df22023e8c219b225b8df4390d10f4e9cc847010b9736a4dc09a51e2affad6afdd99efdee9bc407ba626f94db7ee8222b7a3e51dcf6ec5ffb4f7caa5ce90ea1a825022a0e75e24f475fb7f8f6a454f20c50b4db47064f6857a3db3850767a49af61aaab08e58e1ee6cf6756f05e228ebaa6ee64eeb7b3693dc7b6a47fd9ca06d25b8143d50aed0b543543ce1de500beb8f08e3eef8719eb59c8c7f69bcd887a363ae0bb978acec69ab092534d97bcafd4bbbdb73efb335860689f56c7b0c9be3ad57f6426dc4eb37bd5d7ee78f0c425dca53eaeae6"

type TokenService interface {
	Create(user *types.User, expiresAfter time.Duration) (string, error)
	Read(token string) (jwt.Claims, error)
}

func NewTokenService() TokenService {
	return &JWTService{}
}

type JWTService struct{}

func (js *JWTService) Create(user *types.User, expiresAfter time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.UserClaims{
		Username:  user.Username,
		Email:     user.Email,
		Algorithm: jwt.SigningMethodHS256.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAfter)),
		},
	})

	signedToken, err := token.SignedString(
		[]byte(
			secret,
		),
	)

	if err != nil {
		return "", err
	}
	return signedToken, nil

}

func (js *JWTService) Read(token string) (jwt.Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&types.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("you're Unauthorized due to error parsing the JWT")
			}
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)

	if err != nil || !parsedToken.Valid {
		return nil, fmt.Errorf("you're Unauthorized due to invalid token")
	}

	return parsedToken.Claims, nil
}
