package auth

import (
	"context"
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/shared/logger"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func (s *Service) CreateAccessAuthToken(ctx context.Context, userID int) (*string, error) {

	accessToken, err := generateToken(userID, s.cfg.TokenSecret, s.cfg.AccessTokenTTLMin)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err) // ✅
	}

	logger.Info(ctx, "Access token generated successfully")
	return accessToken, nil
}

func (s *Service) CreateRefreshAuthToken(ctx context.Context, userID int) (*string, error) {

	refreshToken, err := generateToken(userID, s.cfg.TokenSecret, s.cfg.RefreshTokenTTLDays)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	if err := s.SaveRefreshTokenToRedis(ctx, s.cfg.RefreshTokenTTLDays, userID, *refreshToken); err != nil {
		return nil, fmt.Errorf("failed to save refresh token to redis: %w", err)
	}

	logger.Info(ctx, "Refresh token generated and saved successfully")
	return refreshToken, nil

}

func generateToken(userID int, secret string, ttl time.Duration) (*string, error) {
	now := time.Now()
	userIdStr := strconv.Itoa(userID)
	claims := jwt.RegisteredClaims{
		Subject:   userIdStr,
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

func (s *Service) SaveRefreshTokenToRedis(ctx context.Context, ttl time.Duration, userID int, token string) error {

	key := fmt.Sprintf("refresh:userID:%d", userID)
	value := token

	err := s.clnts.RedisClnt.Redis.Set(ctx, key, value, ttl).Err()
	if err != nil {
		logger.Error(ctx, fmt.Errorf("failed to save refresh token to redis: %w", err))
		return err
	}

	logger.Info(ctx, "Refresh token saved to Redis successfully")
	return nil
}
