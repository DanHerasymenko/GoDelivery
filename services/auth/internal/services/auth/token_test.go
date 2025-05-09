package auth

import (
	"context"
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type RedisCommander interface {
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}

type TestService struct {
	clnts struct {
		RedisClnt struct {
			Redis RedisCommander
		}
	}
}

type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	args := m.Called(ctx, keys)

	cmd := redis.NewIntCmd(ctx)
	if err := args.Error(0); err != nil {
		cmd.SetErr(err)
	} else {
		cmd.SetVal(1)
	}
	return cmd
}

func Test_generateToken(t *testing.T) {
	type args struct {
		userID string
		secret string
		ttl    time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				userID: "user",
				secret: "secret",
				ttl:    15 * time.Minute,
			},
			wantErr: false,
		},
		{
			name: "no userID",
			args: args{
				userID: "",
				secret: "secret",
				ttl:    15 * time.Minute,
			},
			wantErr: false,
		},
		{
			name: "no secret",
			args: args{
				userID: "user",
				secret: "",
				ttl:    15 * time.Minute,
			},
			wantErr: false,
		},
		{
			name: "zero ttl",
			args: args{
				userID: "user",
				secret: "secret",
				ttl:    0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := generateToken(tt.args.userID, tt.args.secret, tt.args.ttl)

			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}

func TestService_DeleteTokenFromRedis(t *testing.T) {

	ctx := context.Background()

	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}

	mockRedis := new(MockRedis)
	mockRedis.On("Del", ctx, []string{expectedKey}).Return(nil)

	service := &TestService{}
	service.clnts.RedisClnt.Redis = mockRedis

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				cfg:   tt.fields.cfg,
				clnts: tt.fields.clnts,
			}
			tt.wantErr(t, s.DeleteTokenFromRedis(tt.args.ctx, tt.args.userID), fmt.Sprintf("DeleteTokenFromRedis(%v, %v)", tt.args.ctx, tt.args.userID))
		})
	}
}
