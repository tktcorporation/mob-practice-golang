package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tktcorporation/mob-practice-golang/web-service-gin/album"
	"github.com/tktcorporation/mob-practice-golang/web-service-gin/album/repository"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, repository.GetAll())
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
    var newAlbum album.Album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    repository.Create(newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")

    if a, err := repository.GetById(id); err == nil {
        c.IndentedJSON(http.StatusOK, a)
        return
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}