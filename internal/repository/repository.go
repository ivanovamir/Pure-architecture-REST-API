package repository

import (
	"context"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/pkg/cache"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	errScanRows     = "error occurred scanning rows"
	errParsRows     = "error occurred parsing rows"
	errScanRow      = "error occurred scanning single row"
	errUserTookBook = "error occurred db: user already have this book"
)

type Repository struct {
	Book
	User
}

type Option func(r *Repository)

type Book interface {
	GetAllBooks(ctx context.Context) ([]*dto.Book, error)
	GetBookByID(ctx context.Context, bookID int) (*dto.Book, error)
}

type User interface {
	GetAllUsers(ctx context.Context) ([]*dto.User, error)
	GetUserByID(ctx context.Context, userId int) (*dto.User, error)
	CheckUserBook(ctx context.Context, bookId, userId int) (bool, error)
	TakeBook(ctx context.Context, bookId, userId int) error
	RegisterUser(ctx context.Context, name string) (string, error)
	WriteRefreshToken(ctx context.Context, refreshToken string, userId string, ttl time.Duration) error
	CheckRefreshToken(ctx context.Context, userId string) (string, error)
}

func NewRepository(db *sqlx.DB, cacheClient cache.Cache) *Repository {
	return &Repository{
		Book: NewbookRepository(db),
		User: NewUserRepository(db, cacheClient),
	}
}
