package service

import (
	"context"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/logger"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/password_manager"
)

type userService struct {
	log             *logger.Logger
	repo            repository.User
	passwordManager password_manager.PasswordManager
}

func NewUserService(log *logger.Logger, repo repository.User, passwordManager password_manager.PasswordManager) User {
	return &userService{
		log:             log,
		repo:            repo,
		passwordManager: passwordManager,
	}
}

func (s *userService) Create(ctx context.Context, userDto *dto.User) (string, error) {

	// TODO: Validate user data
	// TODO: Check if user exist
	// TODO: Hash password
	// TODO: Create Refresh Token
	// TODO: Create Access Token

	// Check if user exist
	ok, err := s.repo.IsExist(ctx, userDto.Phone, userDto.Email)

	if err != nil {
		return "", err
	}

	// If user exist -> return err
	if ok {
		return "", ErrUserAlreadyRegistered
	}

	// Hash password
	s.hashPassword(userDto)

	return "", nil
}

func (s *userService) hashPassword(userDto *dto.User) {
	salt := s.passwordManager.Salt()
	userDto.PasswordHash = s.passwordManager.PasswordHash([]byte(userDto.Password), salt)
}
