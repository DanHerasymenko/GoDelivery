package auth

import (
	"context"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/clients/redisClient"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	//TODO implement me
	panic("implement me")
}

func (m *MockRedis) Get(ctx context.Context, key string) *redis.StringCmd {
	//TODO implement me
	panic("implement me")
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

func Test_DeleteTokenFromRedis(t *testing.T) {
	ctx := context.Background()
	userID := "user123"
	expectedKey := "auth:" + userID

	mockRedis := new(MockRedis)
	mockRedis.On("Del", ctx, []string{expectedKey}).Return(nil)

	service := &Service{
		clnts: &clients.Clients{
			RedisClnt: &redisClient.Client{
				Redis: mockRedis, // mock реалізує RedisCommander
			},
		},
	}

	err := service.DeleteTokenFromRedis(ctx, userID)
	assert.NoError(t, err)
	mockRedis.AssertCalled(t, "Del", ctx, []string{expectedKey})
}
