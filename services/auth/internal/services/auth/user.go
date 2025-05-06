package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/constants"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

type User struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	//	Role         string    `json:"role" bson:"role"`
	CreatedAt int64 `json:"created_at"`
}

func (s *Service) CreateUser(ctx context.Context, email, passwordHash string) (*User, error) {

	now := time.Now().Unix()
	var id string
	var pgErr *pgconn.PgError

	user := &User{
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    now,
	}

	query := `
        INSERT INTO users (email, password_hash, created_at)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	err := s.clnts.PostgresClnt.Postgres.QueryRow(ctx, query, email, passwordHash, now).Scan(&id)
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return nil, constants.ErrUserAlreadyExists
	} else if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	user.Id = id

	return user, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*User, error) {

	query := `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`
	user := &User{}

	err := s.clnts.PostgresClnt.Postgres.QueryRow(ctx, query, email).Scan(&user.Id, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, constants.ErrUserNotFound
	}

	return user, err
}
