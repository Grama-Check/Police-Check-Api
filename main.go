package main

import (
	"Police-Check-Api/handler"
	"Police-Check-Api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	authGroup := r.Group("/").Use(middleware.AuthMiddleware())

	authGroup.POST("/", handler.PoliceCheck)

	r.Run(":6060")
}
