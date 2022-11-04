package command

import (
	"context"
	"database/sql"
	"errors"
	"shortener/domain/shortener/model"
)

type urlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) *urlRepository {
	return &urlRepository{db: db}
}

func (s *urlRepository) Add(ctx context.Context, model model.Link) error {
	query := `INSERT INTO "url_shortener" 
    (id, code, full_url, expired_at) 
VALUES ($1, $2, $3, $4);`

	_, err := s.db.ExecContext(ctx, query, model.Id, model.Code, model.FullUrl, model.ExpiredAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *urlRepository) GetByCode(ctx context.Context, code string) (*model.Link, error) {
	query := `SELECT * from "url_shortener" WHERE code = $1`

	result := &model.Link{}
	row := s.db.QueryRowContext(ctx, query, code)

	err := row.Scan(&result.Id, &result.Code, &result.FullUrl, &result.ExpiredAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("no such url")
	case err != nil:
		return result, nil
	default:
		return nil, err
	}
}

func (s *urlRepository) HasByCode(ctx context.Context, code string) (bool, error) {
	query := `SELECT exists(id) FROM "url_shortener" WHERE code = $1`

	var has bool

	row := s.db.QueryRowContext(ctx, query, code)

	err := row.Scan(&has)
	switch {
	case err == sql.ErrNoRows || err == nil:
		return has, nil
	default:
		return false, err
	}
}
