package handler

import (
	"github.com/Futturi/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) Init() *gin.Engine {
	serv := gin.Default()
	api := serv.Group("/url")
	{
		api.POST("/", h.GetLink)
	}
	serv.GET("/:link", h.Redir)
	return serv
}
