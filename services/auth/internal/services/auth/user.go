package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

// business error -> to frontend
var (
	ErrUserAlreadyExists = fmt.Errorf("user with nickname %s already exists")
	ErrUserNotFound      = fmt.Errorf("user not found")
)

type User struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	//	Role         string    `json:"role" bson:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Service) CreateUser(ctx context.Context, email, passwordHash string) (*User, error) {

	now := time.Now()
	var id int
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
		return nil, ErrUserAlreadyExists
	} else if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	user.Id = id

	return user, nil
}
