package repository

import (
	"context"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	User
}

type User interface {
	Create(ctx context.Context, userDto *dto.User) error
	Get(ctx context.Context, id int) (*dto.User, error)
	IsExist(ctx context.Context, phone, email string) (bool, error)
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
