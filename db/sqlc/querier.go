// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, name string) (Author, error)
	CreateBook(ctx context.Context, title string) (Book, error)
	DeleteAuthor(ctx context.Context, id int64) error
	DeleteBook(ctx context.Context, id int64) error
	GetAuthor(ctx context.Context, id int64) (Author, error)
	GetBook(ctx context.Context, id int64) (Book, error)
	ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error)
	ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error)
}

var _ Querier = (*Queries)(nil)
