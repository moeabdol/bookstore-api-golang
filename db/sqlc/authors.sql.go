// Code generated by sqlc. DO NOT EDIT.
// source: authors.sql

package db

import (
	"context"
	"encoding/json"
	"time"
)

const authorIDExists = `-- name: AuthorIDExists :one
SELECT COUNT(*)
FROM authors
WHERE id = $1
`

func (q *Queries) AuthorIDExists(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, authorIDExists, id)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
  name
) VALUES (
  $1
) RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, name)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
WITH author_books AS (
  SELECT id, title, created_at, updated_at, author_id
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
LIMIT 1
`

type GetAuthorParams struct {
	ID     int64 `json:"id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetAuthorRow struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Books     json.RawMessage `json:"books"`
}

func (q *Queries) GetAuthor(ctx context.Context, arg GetAuthorParams) (GetAuthorRow, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, arg.ID, arg.Limit, arg.Offset)
	var i GetAuthorRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Books,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, created_at, updated_at FROM authors
ORDER By name
LIMIT $1
OFFSET $2
`

type ListAuthorsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors
SET name = $2, updated_at = now()
WHERE id = $1
RETURNING id, name, created_at, updated_at
`

type UpdateAuthorParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, updateAuthor, arg.ID, arg.Name)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
