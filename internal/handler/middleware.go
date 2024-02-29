package handler

import (
	"context"
	"net/http"
	"strings"

	proto "github.com/Futturi/AuthSer/protos"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	splitedheader := strings.Split(header, " ")
	if len(splitedheader) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid header"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	userid, err := h.grpcclient.CheckIdentity(context.Background(), &proto.CheckIdentityRequest{Header: splitedheader[1]})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Set("userId", userid)
}
