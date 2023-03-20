package repository

import (
	"context"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]*dto.User, error) {
	var usersDTO []*dto.User

	rows, err := r.db.QueryxContext(ctx, fmt.Sprintf(`SELECT id, name, created_at FROM "user"`))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		userDTO := &dto.User{}
		if err := rows.Scan(
			&userDTO.Id,
			&userDTO.Name,
			&userDTO.CreatedAt,
		); err != nil {
			if err.Error() != "sql: no rows in result set" {
				return nil, err
			}
		}
		usersDTO = append(usersDTO, userDTO)
	}

	return usersDTO, nil
}
