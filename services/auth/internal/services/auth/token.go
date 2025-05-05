package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	mins = time.Minute
	day  = 24 * time.Hour
)

func (s *Service) CreateAccessAuthToken(userID string) (*string, error) {

	accessToken, err := generateToken(userID, s.cfg.TokenSecret, s.cfg.AccessTokenTTLMin*mins)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err) // ✅
	}

	return accessToken, nil
}

func (s *Service) CreateRefreshAuthToken(userID string) (*string, error) {

	refreshToken, err := generateToken(userID, s.cfg.TokenSecret, s.cfg.RefreshTokenTTLDays*day)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	if err := s.SaveRefreshTokenToRedis(*refreshToken); err != nil {
		return nil, fmt.Errorf("failed to save refresh token to redis: %w", err)
	}

	return refreshToken, nil

}

func generateToken(userID string, secret string, ttl time.Duration) (*string, error) {
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

func (s *Service) SaveRefreshTokenToRedis(token string) error {
	return nil
}