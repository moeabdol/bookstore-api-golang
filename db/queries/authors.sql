-- name: CreateAuthor :one
INSERT INTO authors (
  name
) VALUES (
  $1
) RETURNING *;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER By name
LIMIT $1
OFFSET $2;
