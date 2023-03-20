package service

import (
	"context"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
)

type Service struct {
	bookService
	userService
}

type Book interface {
	GetAllBooks(ctx context.Context) ([]*dto.Book, error)
	GetBookByID(ctx context.Context, bookID int) (*dto.Book, error)
}

type User interface {
	GetAllUsers(ctx context.Context) ([]*dto.User, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		bookService: *NewBookService(repo),
		userService: *NewUserService(repo),
	}
}
