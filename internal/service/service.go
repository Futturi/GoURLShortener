package service

import (
	"github.com/Futturi/internal/models"
	"github.com/Futturi/internal/repository"
)

type Service struct {
	Api
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Api: NewApiService(repo.Api)}
}

type Api interface {
	Parse(header string) (int, error)
	GetLink(url models.URL) (string, error)
	Link(link string) (string, error)
}
