package types

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id       string
	Username string
	Email    string
	Password string
}

type Token struct {
	Value string
}

type UserClaims struct {
	UserId    string
	Username  string
	Email     string
	Algorithm string
	jwt.RegisteredClaims
}
