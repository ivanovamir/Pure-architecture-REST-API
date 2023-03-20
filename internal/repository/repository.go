package repository

import (
	"context"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Book
	User
}

type Book interface {
	GetAllBooks(ctx context.Context) ([]*dto.Book, error)
	GetBookByID(ctx context.Context, bookID int) (*dto.Book, error)
}

type User interface {
	GetAllUsers(ctx context.Context) ([]*dto.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Book: NewbookRepository(db),
		User: NewUserRepository(db),
	}
}
