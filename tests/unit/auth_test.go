package unit

import (
	"context"
	"testing"

	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/Alieksieiev0/auth-service/tests/data"
	"github.com/Alieksieiev0/auth-service/tests/mocks"
	"github.com/Alieksieiev0/auth-service/tests/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	m := mocks.NewMockUserServiceClient(gomock.NewController(t))
	utils.ExpectSave(m)

	serv := services.NewAuthService(m, nil)
	err := serv.Register(context.Background(), data.User())
	if err != nil {
		t.Error(err)
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	mu := mocks.NewMockUserServiceClient(ctrl)
	if err := utils.ExpectGetByUsername(mu); err != nil {
		t.Error(err)
	}

	mt := mocks.NewMockTokenService(ctrl)
	utils.ExpectCreate(mt)

	serv := services.NewAuthService(mu, mt)
	result, err := serv.Login(context.Background(), data.User())
	if err != nil {
		t.Error(err)
	}

	expected := data.UserToken()
	assert.Equal(t, expected.User, result.User, "Login returned unexpected result")
	assert.Equal(t, expected.Token, result.Token, "Login returned unexpected result")
}

func TestReadClaims(t *testing.T) {
	m := mocks.NewMockTokenService(gomock.NewController(t))
	utils.ExpectRead(m)
	serv := services.NewAuthService(nil, m)
	result, err := serv.ReadClaims(context.Background(), data.Token())
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, data.UserClaims(), result, "ReadClaims return invalid claims")
}
