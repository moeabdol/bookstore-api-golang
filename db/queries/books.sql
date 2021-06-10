-- name: CreateBook :one
INSERT INTO books (
  title
) VALUES (
  $1
) RETURNING *;
