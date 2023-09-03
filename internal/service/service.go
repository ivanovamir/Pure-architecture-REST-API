package service

import (
	"context"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/logger"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/password_manager"
)

type Service struct {
	User
}

type User interface {
	Create(ctx context.Context, userDto *dto.User) (string, error)
}

func NewService(log *logger.Logger, repo *repository.Repository, passwordManager password_manager.PasswordManager) *Service {
	userService := NewUserService(log, repo, passwordManager)
	return &Service{
		User: userService,
	}
}
