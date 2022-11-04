package command

import (
	"context"
	"shortener/domain/shortener/model"
)

type LinkRepository interface {
	Add(ctx context.Context, model model.Link) error
	GetByCode(ctx context.Context, code string) (*model.Link, error)
	HasByCode(ctx context.Context, code string) (bool, error)
}

type CacheStore interface {
	Add(ctx context.Context, model model.Link) error
	GetByCode(ctx context.Context, code string) (*model.Link, error)
	HasByCode(ctx context.Context, code string) (bool, error)
}
