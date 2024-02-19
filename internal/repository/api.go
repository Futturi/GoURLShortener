package repository

import (
	"context"
	"errors"

	"github.com/Futturi/internal/models"
	"github.com/redis/go-redis/v9"
)

type RepositoryApi struct {
	db *redis.Client
}

func NewRepositoryApi(db *redis.Client) *RepositoryApi {
	return &RepositoryApi{db: db}
}

func (r *RepositoryApi) GetLink(url models.URL, newlink string) error {
	ctx := context.Background()
	err := r.db.Set(ctx, url.Url, newlink, 0)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}

func (r *RepositoryApi) Link(link string) (string, error) {
	ctx := context.Background()
	entity := r.db.Get(ctx, link)
	if entity.Err() == redis.Nil {
		return "", errors.New("your link not in db")
	}
	if entity.Err() != nil {
		return "", entity.Err()
	}
	return entity.String(), nil
}
