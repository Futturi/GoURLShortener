package repository

import (
	"github.com/Futturi/internal/models"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Api
}

func NewRepository(db *redis.Client) *Repository {
	return &Repository{Api: NewRepositoryApi(db)}
}

type Api interface {
	GetLink(url models.URL, newlink string) error
	Link(link string) (string, error)
}
