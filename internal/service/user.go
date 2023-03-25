package service

import (
	"context"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
)

type userService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *userService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]*dto.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, userId int) (*dto.User, error) {
	return s.repo.GetUserByID(ctx, userId)
}

func (s *userService) TakeBook(ctx context.Context, bookId, userId int) error {
	if bookId == 0 || userId == 0 {
		return fmt.Errorf("")
	}

	ok, err := s.repo.CheckUserBook(ctx, bookId, userId)

	if !ok {
		return s.repo.TakeBook(ctx, bookId, userId)
	}
	return err
}
