package handler

import (
	"net/http"

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
	c.Redirect(http.StatusPermanentRedirect, result)
}
