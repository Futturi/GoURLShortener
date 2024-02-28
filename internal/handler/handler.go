package handler

import (
	proto "github.com/Futturi/AuthSer/protos"
	"github.com/Futturi/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service    *service.Service
	grpcclient proto.AuthClient
}

func NewHandler(service *service.Service, grpcclient proto.AuthClient) *Handler {
	return &Handler{service: service, grpcclient: grpcclient}
}
func (h *Handler) Init() *gin.Engine {
	serv := gin.Default()
	auth := serv.Group("/auth")
	{
		auth.POST("/signup", h.Register)
		auth.POST("/signin", h.Login)
	}
	api := serv.Group("/url", h.CheckIdentity)
	{
		api.POST("/", h.GetLink)
	}
	serv.GET("/:link", h.Redir)
	return serv
}
