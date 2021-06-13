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
WITH author_books AS (
  SELECT *
  FROM books
  WHERE author_id = $1
  ORDER BY id
  LIMIT $2
  OFFSET $3
)
SELECT authors.id, authors.name, authors.created_at, authors.updated_at, json_agg(author_books.*) as books
FROM authors
LEFT JOIN author_books
ON author_books.author_id = authors.id
WHERE authors.id = $1
GROUP BY authors.id
LIMIT 1;

-- name: UpdateAuthor :one
UPDATE authors
SET name = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: AuthorIDExists :one
SELECT COUNT(*)
FROM authors
WHERE id = $1;
