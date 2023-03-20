package repository

import (
	"context"
	"fmt"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/jmoiron/sqlx"
)

type bookRepository struct {
	db *sqlx.DB
}

func NewbookRepository(db *sqlx.DB) *bookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAllBooks(ctx context.Context) ([]*dto.Book, error) {
	var booksDTO []*dto.Book

	rows, err := r.db.QueryxContext(ctx, fmt.Sprintf(`
		SELECT
    	book.id,
    	book.title,
    	book.year,
    	a.id,
    	a.Name,
    	g.id,
    	g.title
		FROM book
    	     INNER JOIN author a ON a.id = book.author_id
    	     INNER JOIN genre g ON g.id = book.genre_id;`))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		bookDTO := &dto.Book{}
		if err = rows.Scan(
			&bookDTO.Id,
			&bookDTO.Title,
			&bookDTO.Year,
			&bookDTO.Author.Id,
			&bookDTO.Author.Name,
			&bookDTO.Genre.Id,
			&bookDTO.Genre.Title,
		); err != nil {
			return nil, err
		}

		booksDTO = append(booksDTO, bookDTO)

		if err := rows.Err(); err != nil {
			if err.Error() != "sql: no rows in result set" {
				return nil, err
			}
		}
	}

	return booksDTO, nil
}

func (r *bookRepository) GetBookByID(ctx context.Context, bookID int) (*dto.Book, error) {
	bookDTO := &dto.Book{}

	row := r.db.QueryRowxContext(ctx, fmt.Sprintf(`
		SELECT
    	book.id,
    	book.title,
    	book.year,
    	a.id,
    	a.Name,
		g.id,
    	g.title
    	FROM book
    	    INNER JOIN genre g ON g.id = book.genre_id
    	    INNER JOIN author a on a.id = book.author_id
    			WHERE book.id = %d LIMIT 1`, bookID))

	if row.Err() != nil {
		return nil, row.Err()
	}

	if err := row.Scan(
		&bookDTO.Id,
		&bookDTO.Title,
		&bookDTO.Year,
		&bookDTO.Author.Id,
		&bookDTO.Author.Name,
		&bookDTO.Genre.Id,
		&bookDTO.Genre.Title,
	); err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		}
	}

	return bookDTO, nil
}
