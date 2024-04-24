package unit

import (
	"testing"

	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/Alieksieiev0/auth-service/internal/types"
	"github.com/Alieksieiev0/auth-service/tests/data"
	"github.com/Alieksieiev0/auth-service/tests/utils"
	"github.com/stretchr/testify/assert"
)

var serv = services.NewTokenService()

func TestCreate(t *testing.T) {
	token, err := serv.Create(data.UserPublic(), utils.Exp)
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Errorf("Token was not created")
	}
}

func TestRead(t *testing.T) {
	assert := assert.New(t)
	user := data.UserPublic()

	token, err := serv.Create(user, utils.Exp)
	if err != nil {
		t.Error(err)
	}

	claims, err := serv.Read(token)
	if err != nil {
		t.Error(err)
	}

	userClaims, ok := claims.(*types.UserClaims)
	if !ok {
		t.Errorf("Wrong claims were generated")
	}

	assert.Equal(user.Id, userClaims.UserId, "Wrong User Id in claims")
	assert.Equal(user.Email, userClaims.Email, "Wrong User Email in claims")
	assert.Equal(user.Username, userClaims.Username, "Wrong User Name in claims")
}
