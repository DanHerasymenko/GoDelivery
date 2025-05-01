package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GeneratePasswordHash(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate password hash: %w", err)
	}

	return string(res), nil
}

func (s *Service) ComparePasswordHashAndPassword(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, fmt.Errorf("failed to compare password hash and password: %w", err)
	}

	return true, nil
}
