-- name: CreateBook :one
INSERT INTO books (
  title,
  author_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: ListBooks :many
SELECT books.id, books.title, books.created_at, books.updated_at, authors.id as author_id, authors.name as author_name
FROM books
INNER JOIN authors
ON books.author_id = authors.id
ORDER BY books.created_at
LIMIT $1
OFFSET $2;

-- name: GetBook :one
SELECT books.id, books.title, books.created_at, books.updated_at, authors.id as author_id, authors.name as author_name
FROM books
INNER JOIN authors
ON books.author_id = authors.id
WHERE books.id = $1
LIMIT 1;

-- name: UpdateBook :one
UPDATE books
SET title = $2, author_id = $3, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: BookTitleExists :one
SELECT COUNT(*)
FROM books
WHERE title = $1;
