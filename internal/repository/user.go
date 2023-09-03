package repository

import (
	"context"
	"errors"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/repository/dto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) User {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, userDto *dto.User) error {
	q := `INSERT INTO "user" (name, surname, patronymic, password_hash) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(ctx, q, userDto.Name, userDto.Surname, userDto.Patronymic, userDto.PasswordHash).Scan(&userDto.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotCreated
		}
		return err
	}
	return nil
}

func (r *userRepository) Get(ctx context.Context, id int) (*dto.User, error) {
	userDto := new(dto.User)
	q := `SELECT u.id, u.name, u.surname, u.patronymic FROM "user" u WHERE u.id = $1`

	err := r.db.QueryRow(ctx, q, id).Scan(&userDto.Id, &userDto.Name, &userDto.Surname, &userDto.Patronymic)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return userDto, nil
}

func (r *userRepository) IsExist(ctx context.Context, phone, email string) (bool, error) {
	var isExist bool
	q := `SELECT EXISTS(SELECT 1 FROM "user" u WHERE u.phone = $1 OR u.email = $2)`

	err := r.db.QueryRow(ctx, q, phone, email).Scan(&isExist)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return isExist, err
	}

	return isExist, nil
}
