package service

import (
	"context"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/token_manager"
)

type userService struct {
	repo         repository.User
	tokenManager token_manager.TokenManager
}

func NewUserService(repo repository.User, tokenManager token_manager.TokenManager) *userService {
	return &userService{repo: repo, tokenManager: tokenManager}
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

func (s *userService) RegisterUser(ctx context.Context, name string) (*dto.SuccessRegister, error) {
	userId, err := s.repo.RegisterUser(ctx, name)

	if err != nil {
		return nil, err
	}

	accessToken, err := s.tokenManager.NewJWT(userId)

	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenManager.NewRefreshToken()

	if err != nil {
		return nil, err
	}

	if err = s.repo.WriteRefreshToken(ctx, refreshToken, userId); err != nil {
		return nil, err
	}

	return &dto.SuccessRegister{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
