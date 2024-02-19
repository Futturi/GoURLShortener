package service

import (
	"fmt"
	"math/rand"

	"github.com/Futturi/internal/models"
	"github.com/Futturi/internal/repository"
)

const (
	salt = "rgjoweifAWE:Oifse;orthijgerjng;o3ij;oier2ueheaoijrtoi"
)

type ApiService struct {
	repo repository.Api
}

func NewApiService(repo repository.Api) *ApiService {
	return &ApiService{repo: repo}
}
func (a *ApiService) GetLink(url models.URL) (string, error) {
	newlink := GenerateNewLink(url)
	return newlink, a.repo.GetLink(url, newlink)
}

func GenerateNewLink(url models.URL) string {
	result := make([]byte, 10)
	rand.Read(result)
	return fmt.Sprintf("%x", result)
}

func (a *ApiService) Link(link string) (string, error) {
	return a.repo.Link(link)
}
