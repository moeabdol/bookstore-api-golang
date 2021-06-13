// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  password,
  email
) VALUES (
  $1, $2, $3
) RETURNING id, username, password, email, created_at, updated_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const emailExists = `-- name: EmailExists :one
SELECT COUNT(*)
FROM users
WHERE email = $1
`

func (q *Queries) EmailExists(ctx context.Context, email string) (int64, error) {
	row := q.db.QueryRowContext(ctx, emailExists, email)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, created_at, updated_at
FROM users
WHERE id = $1
LIMIT 1
`

type GetUserRow struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) GetUser(ctx context.Context, id int64) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}