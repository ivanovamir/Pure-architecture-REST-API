package service

import (
	"context"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
)

type Service struct {
	Book
	User
}

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Book interface {
	GetAllBooks(ctx context.Context) ([]*dto.Book, error)
	GetBookByID(ctx context.Context, bookID int) (*dto.Book, error)
}

type User interface {
	GetAllUsers(ctx context.Context) ([]*dto.User, error)
	GetUserByID(ctx context.Context, userId int) (*dto.User, error)
	TakeBook(ctx context.Context, bookId, userId int) error
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repo),
		User: NewUserService(repo),
	}
}
