package types

import "github.com/golang-jwt/jwt/v5"

type UserToken struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Value string `json:"value"`
}

type UserClaims struct {
	UserId    string `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Algorithm string `json:"algorithm"`
	jwt.RegisteredClaims
}
