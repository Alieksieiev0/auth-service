package utils

import (
	"time"

	"github.com/Alieksieiev0/auth-service/tests/data"
	"github.com/Alieksieiev0/auth-service/tests/mocks"
	"github.com/golang/mock/gomock"
)

const Exp time.Duration = time.Hour * 6

func ExpectGetByUsername(m *mocks.MockUserServiceClient) error {
	user := data.User()
	username := data.ProtoUsernameRequest(user)
	userProto, err := data.ProtoUserResponse(user)
	if err != nil {
		return err
	}

	m.EXPECT().
		GetByUsername(gomock.Any(), gomock.Eq(username)).
		Return(userProto, nil)
	return nil
}

func ExpectSave(m *mocks.MockUserServiceClient) {
	userProto := data.ProtoUserRequest(data.User())
	m.EXPECT().Save(gomock.Any(), gomock.Eq(userProto)).Return(nil, nil)
}

func ExpectCreate(m *mocks.MockTokenService) {
	userPublic := data.UserPublic()
	token := data.Token()

	m.EXPECT().Create(gomock.Eq(userPublic), gomock.Eq(Exp)).Return(token.Value, nil)
}

func ExpectRead(m *mocks.MockTokenService) {
	token := data.Token()
	claims := data.UserClaims()
	m.EXPECT().Read(gomock.Eq(token.Value)).Return(claims, nil)
}
