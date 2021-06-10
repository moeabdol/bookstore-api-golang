-- name: CreateBook :one
INSERT INTO books (
  title
) VALUES (
  $1
) RETURNING *;


-- name: ListBooks :many
SELECT * FROM books
ORDER BY created_at
LIMIT $1
OFFSET $2;
