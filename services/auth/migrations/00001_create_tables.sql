-- +goose Up
CREATE TABLE IF NOT EXISTS users (
                                       id SERIAL PRIMARY KEY,
                                       email VARCHAR(255) UNIQUE NOT NULL,
                                       password_hash DOUBLE PRECISION,
                                       created_at BIGINT NOT NULL,
    );

-- +goose Down
DROP TABLE IF EXISTS users;


--- goose -dir ./migrations postgres "postgres://user:password@localhost:port/dbname?sslmode=disable" up
