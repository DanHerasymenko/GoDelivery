-- +goose Up
CREATE TABLE IF NOT EXISTS users (
                                       id SERIAL PRIMARY KEY,
                                       email VARCHAR(255) UNIQUE NOT NULL,
                                       password_hash TEXT NOT NULL,
                                       created_at BIGINT NOT NULL
    );

-- +goose Down
DROP TABLE IF EXISTS users;


--- goose -dir ./migrations postgresClient "postgresClient://user:password@localhost:port/dbname?sslmode=disable" up
