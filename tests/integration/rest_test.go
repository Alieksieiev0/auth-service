package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/Alieksieiev0/auth-service/internal/transport/rest"
	"github.com/Alieksieiev0/auth-service/internal/types"
	"github.com/Alieksieiev0/auth-service/tests/data"
	"github.com/Alieksieiev0/auth-service/tests/mocks"
	"github.com/Alieksieiev0/auth-service/tests/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRESTServer(t *testing.T) {
	assert := assert.New(t)
	user := data.User()
	ctrl := gomock.NewController(t)
	mu, mt, err := prepareServices(ctrl)
	mu.EXPECT().
		Save(gomock.Any(), gomock.Not(gomock.Eq(data.ProtoUserRequest(user)))).
		Return(nil, nil)
	if err != nil {
		t.Error(err)
	}

	serv := services.NewAuthService(mu, mt)
	app := fiber.New()
	s := rest.NewServer(app, ":3011")
	go func() {
		if err = s.Start(serv); err != nil {
			t.Error(err)
		}
	}()

	body, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	err = register(bytes.NewReader(body), assert)
	if err != nil {
		t.Error(err)
	}

	err = login(bytes.NewReader(body), assert)
	if err != nil {
		t.Error(err)
	}
}

func register(body io.Reader, assert *assert.Assertions) error {
	resp, err := http.Post("http://localhost:3011/register", "application/json", body)
	if err != nil {
		return err
	}

	assert.Equal(
		http.StatusCreated,
		resp.StatusCode,
		"Register endpoint returned incorrect status code",
	)
	return nil
}

func login(body io.Reader, assert *assert.Assertions) error {
	resp, err := http.Post("http://localhost:3011/login", "application/json", body)
	if err != nil {
		return err
	}

	assert.Equal(http.StatusOK, resp.StatusCode, "Login endpoint returned incorrect status code")

	respBody := &types.UserToken{}
	err = json.NewDecoder(resp.Body).Decode(respBody)
	if err != nil {
		return fmt.Errorf("Error parsing login endpoint response: %s", err.Error())
	}

	token := data.Token()
	assert.Equal(token.Value, respBody.Token.Value, "Login endpoint returned incorrect token value")
	return nil
}

func prepareServices(
	ctrl *gomock.Controller,
) (*mocks.MockUserServiceClient, *mocks.MockTokenService, error) {
	mu := mocks.NewMockUserServiceClient(ctrl)
	if err := utils.ExpectGetByUsername(mu); err != nil {
		return nil, nil, err
	}

	mt := mocks.NewMockTokenService(ctrl)
	utils.ExpectCreate(mt)

	return mu, mt, nil
}
