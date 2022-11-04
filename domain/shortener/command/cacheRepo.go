package command

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"shortener/domain/shortener/model"
	"time"
)

type cacheRepository struct {
	rdb *redis.Client
}

func NewCacheRepository(rdb *redis.Client) *cacheRepository {
	return &cacheRepository{rdb: rdb}
}

func (s *cacheRepository) Add(ctx context.Context, model model.Link) error {
	dataBytes, err := json.Marshal(&model)
	if err != nil {
		return err
	}
	if err := s.rdb.Set(ctx, model.Code, dataBytes, 2*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func (s *cacheRepository) GetByCode(ctx context.Context, code string) (*model.Link, error) {
	val, err := s.rdb.Get(ctx, code).Result()
	if err != nil {
		return nil, err
	}

	result := &model.Link{}
	if err := json.Unmarshal([]byte(val), result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *cacheRepository) HasByCode(ctx context.Context, code string) (bool, error) {
	_, err := s.rdb.Get(ctx, code).Result()
	switch {
	case err == redis.Nil:
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}
