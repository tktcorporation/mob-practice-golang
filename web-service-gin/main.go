package main

import (
	"github.com/gin-gonic/gin"

	"github.com/tktcorporation/mob-practice-golang/web-service-gin/album/service"
)

func main() {
    router := gin.Default()
    router.GET("/albums", service.GetAlbums)
	router.POST("/albums", service.PostAlbums)
	router.GET("/albums/:id", service.GetAlbumByID)

    router.Run("localhost:8080")
}