package repository

import (
	"context"
	"errors"
	"log"

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
	err := r.db.Set(ctx, newlink, url.Url, 0)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}

func (r *RepositoryApi) Link(link string) (string, error) {
	var link1 string
	ctx := context.Background()
	entity := r.db.Get(ctx, link)
	if err := entity.Scan(&link1); err != nil {
		return "", err
	}
	if entity.Err() == redis.Nil {
		return "", errors.New("your link not in db")
	}
	if entity.Err() != nil {
		return "", entity.Err()
	}
	log.Println(link1)
	return link1, nil
}
