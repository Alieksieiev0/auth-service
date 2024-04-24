package integration

import (
	"context"
	"testing"

	"github.com/Alieksieiev0/auth-service/api/proto"
	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/Alieksieiev0/auth-service/internal/transport/grpc"
	"github.com/Alieksieiev0/auth-service/tests/data"
	"github.com/Alieksieiev0/auth-service/tests/mocks"
	"github.com/Alieksieiev0/auth-service/tests/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const GRPCAddr = ":3012"

func TestGRPCServer(t *testing.T) {
	assert := assert.New(t)
	m := mocks.NewMockTokenService(gomock.NewController(t))
	utils.ExpectRead(m)

	s := grpc.NewServer(GRPCAddr)
	go func() {
		if err := s.Start(services.NewAuthService(nil, m)); err != nil {
			t.Error(err)
		}
	}()

	conn, err := grpc.NewConnection(GRPCAddr)
	if err != nil {
		t.Error(err)
	}
	c := proto.NewAuthServiceClient(conn)

	claims, err := c.ReadClaims(context.Background(), data.ProtoToken())
	if err != nil {
		t.Error(err)
	}
	user := data.UserPublic()
	assert.Equal(user.Id, claims.UserId, "Wrong User Id in gRPC returned user claims")
	assert.Equal(user.Username, claims.Username, "Wrong User Name in gRPC returned user claims")
	assert.Equal(user.Email, claims.Email, "Wrong User Email in gRPC returned user claims")

	if err = s.Stop(); err != nil {
		t.Error(err)
	}
}
