package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	})

	r.Use(corsConfig)

	// health check
	r.GET("/")
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	return r
}
