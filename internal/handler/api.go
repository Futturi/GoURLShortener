package handler

import (
	"context"
	"fmt"
	"net/http"

	proto "github.com/Futturi/AuthSer/protos"
	"github.com/Futturi/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetLink(c *gin.Context) {
	var url models.URL
	err := c.BindJSON(&url)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	result, err := h.service.GetLink(url)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"new_url": result,
	})
}

func (h *Handler) Redir(c *gin.Context) {
	link := c.Param("link")
	result, err := h.service.Link(link)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	fmt.Println(result)
	c.Redirect(http.StatusMovedPermanently, result)
}

func (h *Handler) Register(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	usergrpc := proto.RegisterRequest{Email: user.Email, Password: user.Password}
	response, err := h.grpcclient.Register(context.Background(), &usergrpc)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": response.UserId,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	usergrpc := proto.LoginRequest{Email: user.Email, Password: user.Password}
	token, err := h.grpcclient.Login(context.Background(), &usergrpc)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token.Token,
	})
}
