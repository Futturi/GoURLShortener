package service

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Futturi/internal/models"
	"github.com/Futturi/internal/repository"
	"github.com/golang-jwt/jwt"
)

const (
	salt   = "rgjoweifAWE:Oifse;orthijgerjng;o3ij;oier2ueheaoijrtoi"
	jwtKey = "eojgnrwijnweijfweijfnweijfniwjenfiwnsiquw"
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
	if len(strings.Split(link, "://")) == 1 {
		rlink, err := a.repo.Link(link)
		if err != nil {
			return "", err
		}
		newlin := "https://" + rlink
		return newlin, nil
	} else {
		rlink, err := a.repo.Link(strings.Split(link, "://")[1])
		if err != nil {
			return "", err
		}
		return "https://" + rlink, err
	}
}

type ClaimsUser struct {
	Id int
	jwt.StandardClaims
}

func (a *ApiService) Parse(header string) (int, error) {
	token, err := jwt.ParseWithClaims(header, &ClaimsUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return 0, err
	}
	Claims, ok := token.Claims.(*ClaimsUser)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return Claims.Id, nil
}
