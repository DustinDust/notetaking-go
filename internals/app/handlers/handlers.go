package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
}

func (h *Handlers) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
