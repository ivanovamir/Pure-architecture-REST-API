package service

import (
	"context"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
)

type bookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *bookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks(ctx context.Context) ([]*dto.Book, error) {
	return s.repo.GetAllBooks(ctx)
}

func (s *bookService) GetBookByID(ctx context.Context, bookID int) (*dto.Book, error) {
	return s.repo.GetBookByID(ctx, bookID)
}
