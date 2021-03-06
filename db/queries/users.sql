-- name: CreateUser :one
INSERT INTO users (
  username,
  password,
  email
) VALUES (
  $1, $2, $3
) RETURNING id, username, email, created_at, updated_at;

-- name: GetUser :one
SELECT id, username, email, created_at, updated_at
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: EmailExists :one
SELECT COUNT(*)
FROM users
WHERE email = $1;
