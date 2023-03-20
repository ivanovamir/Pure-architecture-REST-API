package service

import (
	"context"
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
