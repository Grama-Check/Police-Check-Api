package main

import (
	"Police-Check-Api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authGroup := r.Group("/")

	authGroup.POST("/", handler.PoliceCheck)

	r.Run("6060")
}
