package get

import (
	"context"
	"shortener/domain/shortener/command"
	"shortener/domain/shortener/model"
)

type handler struct {
	linkRepo  command.LinkRepository
	cacheRepo command.CacheStore
}

func NewHandler(linkRepo command.LinkRepository, cacheRepo command.CacheStore) *handler {
	return &handler{
		linkRepo:  linkRepo,
		cacheRepo: cacheRepo,
	}
}

func (h *handler) Handle(ctx context.Context, code string) (*model.Link, error) {
	hasCache, err := h.cacheRepo.HasByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if hasCache {
		return h.cacheRepo.GetByCode(ctx, code)
	}

	return h.linkRepo.GetByCode(ctx, code)
}
