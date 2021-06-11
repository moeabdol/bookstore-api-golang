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

-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1
LIMIT 1;
