-- name: CreateAuthor :one
INSERT INTO authors (
  name
) VALUES (
  $1
) RETURNING *;
