package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maminataei/car_seller/api/routers"
	"github.com/maminataei/car_seller/config"
)

func InitServer() {
	cng := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		healthRouterGroup := v1.Group("/health")
		routers.HealthRouter(healthRouterGroup)
	}

	server := &http.Server{
		Handler:     r,
		Addr:        fmt.Sprintf("0.0.0.0:%s", cng.Server.Port),
		ReadTimeout: time.Second * 10,
	}
	server.ListenAndServe()
	fmt.Printf("Server is running on port %s", cng.Server.Port)
}
