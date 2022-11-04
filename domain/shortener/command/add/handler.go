package add

import (
	"context"
	"shortener/domain/shortener/command"
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

func (h *handler) Handle(ctx context.Context, url string) (string, error) {
	model := command.NewLink(url)
	if err := h.linkRepo.Add(ctx, *model.Link); err != nil {
		return "", err
	}

	if err := h.cacheRepo.Add(ctx, *model.Link); err != nil {
		return "", err
	}

	return model.Link.Code, nil
}
