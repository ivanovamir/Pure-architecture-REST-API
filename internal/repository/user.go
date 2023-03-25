package repository

import (
	"context"
	"database/sql"
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
		return nil, fmt.Errorf("%s", errParsRows)
	}

	defer rows.Close()

	for rows.Next() {
		userDTO := &dto.User{}
		if err := rows.Scan(
			&userDTO.Id,
			&userDTO.Name,
			&userDTO.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("%s", errScanRows)
		}
		usersDTO = append(usersDTO, userDTO)
	}

	return usersDTO, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userId int) (*dto.User, error) {
	var userDTO *dto.User

	rows, err := r.db.QueryxContext(ctx, fmt.Sprintf(`
		SELECT 
		"user".id,
		"user".name,
		"user".created_at,
		b.id,
		b.title,
		b.year as created_at,
		g.id,
		g.title,
		a.id,
		a.Name
		FROM "user"
			INNER JOIN user_book ub ON "user".id = ub.user_id
			INNER JOIN book b ON b.id = ub.book_id
			INNER JOIN author a ON a.id = b.author_id
			INNER JOIN genre g ON g.id = b.genre_id
				WHERE "user".id = 1;`))

	if err != nil {
		return nil, fmt.Errorf("%s", errParsRows)
	}

	var userMap bool
	var bookMap = map[string]struct{}{}

	for rows.Next() {
		userDTO := &dto.User{}
		bookDTO := &dto.Book{}

		err := rows.StructScan(&userDTO)
		err = rows.StructScan(&userDTO)
		err = rows.StructScan(&userDTO)
		err = rows.StructScan(&userDTO)

		if err != nil {
			return nil, fmt.Errorf("%s", errScanRows)
		}

		if !userMap {
			userDTO = userDTO
			userMap = true
		}

		_, ok := bookMap[bookDTO.Id]

		if !ok {
			userDTO.Books = append(userDTO.Books, *bookDTO)
		}
	}
	return userDTO, nil

}

func (r *userRepository) TakeBook(ctx context.Context, bookId, userId int) error {

	var userBook struct {
		UserId int
		BookId int
	}

	row := r.db.QueryRowxContext(ctx, fmt.Sprintf(`SELECT * FROM user_book where user_id = %d and book_id = %d`, userId, bookId))

	if err := row.Scan(&userBook.UserId, &userBook.BookId); err != nil {
		if err == sql.ErrNoRows {
			result, err := r.db.ExecContext(ctx, fmt.Sprintf(`INSERT INTO user_book VALUES (%d,%d)`, userId, bookId))
			if err != nil {
				return err
			}

			rows, err := result.RowsAffected()

			if err != nil {
				return err
			}

			if rows != 1 {
				return err
			}
			return nil
		}
		return fmt.Errorf("%s", errScanRow)
	} else {
		return fmt.Errorf("%s", errUserTookBook)

	}
}
