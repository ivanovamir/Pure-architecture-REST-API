package service

import (
	"context"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/token_manager"
	"time"
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
	RegisterUser(ctx context.Context, name string) (*dto.SuccessRegister, error)
	UserValidate(accessToken string) (string, error)
	GetSubFromToken(accessToken string) (string, error)
	UpdateAccessToken(ctx context.Context, accessToken, refreshToken string) (string, error)
	UpdateRefreshToken(ctx context.Context, accessToken, refreshToken string) (*dto.UpdateTokens, error)
}

func NewService(repo *repository.Repository, tokenManager token_manager.TokenManager, refreshTokenTtl time.Duration) *Service {
	return &Service{
		Book: NewBookService(repo),
		User: NewUserService(repo, tokenManager, refreshTokenTtl),
	}
}
