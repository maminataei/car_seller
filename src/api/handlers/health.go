package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() HealthHandler {
	return HealthHandler{}
}

func (h HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world!")
}

func (h HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func (h HealthHandler) HealthPostId(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, "pong by "+id)
}
