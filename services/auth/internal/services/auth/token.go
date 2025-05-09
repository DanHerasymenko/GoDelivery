package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/shared/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

func (s *Service) CreateAccessAuthToken(ctx context.Context, userID string) (*string, error) {

	accessToken, err := generateToken(userID, s.cfg.TokenSecret, s.cfg.AccessTokenTTLMin)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err) // ✅
	}

	logger.Info(ctx, "Access token generated successfully")
	return accessToken, nil
}

func (s *Service) CreateRefreshAuthToken(ctx context.Context, userID string) (*string, error) {

	refreshToken, err := generateToken(userID, s.cfg.TokenSecret, s.cfg.RefreshTokenTTLDays)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	if err := s.saveNewRefreshTokenToRedis(ctx, s.cfg.RefreshTokenTTLDays, userID, *refreshToken); err != nil {
		return nil, fmt.Errorf("failed to save refresh token to redis: %w", err)
	}

	logger.Info(ctx, "Refresh token generated and saved successfully")
	return refreshToken, nil

}

func generateToken(userID string, secret string, ttl time.Duration) (*string, error) {

	if ttl == 0 {
		return nil, errors.New("token ttl must be greater than zero")
	}

	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %w", err)
	}

	return &signedToken, nil
}

func (s *Service) VerifyJWTToken(tokenStr string) (string, error) {

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	tokenStr = strings.TrimSpace(tokenStr)

	claims := jwt.RegisteredClaims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.TokenSecret), nil
	})
	if err != nil || !tkn.Valid {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	// get userID from token subject
	return claims.Subject, nil

}

func (s *Service) saveNewRefreshTokenToRedis(ctx context.Context, ttl time.Duration, userID string, token string) error {

	key := fmt.Sprintf("refresh:userID:%s", userID)
	value := token

	err := s.clnts.RedisClnt.Redis.Set(ctx, key, value, ttl).Err()
	if err != nil {
		logger.Error(ctx, fmt.Errorf("failed to save refresh token to redis: %w", err))
		return err
	}

	logger.Info(ctx, "Refresh token saved to Redis successfully")
	return nil
}

func (s *Service) IfRefreshTokenExistsInRedis(ctx context.Context, token, userID string) (bool, error) {
	key := fmt.Sprintf("refresh:userID:%s", userID)

	val, err := s.clnts.RedisClnt.Redis.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Info(ctx, "Refresh token not found in Redis")
			return false, nil
		}
		logger.Error(ctx, fmt.Errorf("failed to check refresh token in redis: %w", err))
		return false, err
	}

	return val == token, nil
}

func (s *Service) DeleteTokenFromRedis(ctx context.Context, userID string) error {
	
	key := fmt.Sprintf("refresh:userID:%s", userID)
	return s.clnts.RedisClnt.Redis.Del(ctx, key).Err()
}
