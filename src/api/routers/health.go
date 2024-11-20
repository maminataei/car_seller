package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maminataei/car_seller/api/handlers"
)

func HealthRouter(r *gin.RouterGroup) {
	healthHandler := handlers.NewHealthHandler()
	r.GET("/health", healthHandler.Health)
	r.POST("/health", healthHandler.HealthPost)
	r.POST("/:id", healthHandler.HealthPostId)
}
