package main

import (
	"example/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.PostAlbums)

	router.Run("localhost:8080")
}
