// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateBook(ctx context.Context, title string) (Book, error)
	GetBook(ctx context.Context, id int64) (Book, error)
	ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error)
}

var _ Querier = (*Queries)(nil)
