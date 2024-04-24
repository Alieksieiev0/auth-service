package data

import (
	"github.com/Alieksieiev0/auth-service/api/proto"
	"github.com/Alieksieiev0/auth-service/internal/types"
	"golang.org/x/crypto/bcrypt"
)

func User() *types.User {
	return &types.User{
		Username: "test",
		Email:    "test@test.com",
		Password: "something",
	}
}

func UserPublic() *types.User {
	u := User()
	u.Id = "1234"
	u.Password = ""

	return u
}

func UserClaims() *types.UserClaims {
	return &types.UserClaims{
		UserId:    "1234",
		Username:  "test",
		Email:     "test@test.com",
		Algorithm: "some",
	}
}

func ProtoUserRequest(user *types.User) *proto.UserRequest {
	return &proto.UserRequest{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ProtoUserResponse(user *types.User) (*proto.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	userProto := &proto.UserResponse{
		Id:       "1234",
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	}
	return userProto, nil
}

func ProtoUsernameRequest(user *types.User) *proto.UsernameRequest {
	return &proto.UsernameRequest{
		Username: user.Username,
	}
}

func Token() *types.Token {
	return &types.Token{
		Value: "token",
	}
}

func ProtoToken() *proto.Token {
	return &proto.Token{
		Value: "token",
	}
}

func UserToken() *types.UserToken {
	return &types.UserToken{
		User:  *UserPublic(),
		Token: *Token(),
	}
}
