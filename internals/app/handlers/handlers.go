package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
