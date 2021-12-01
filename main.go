package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suren-m/go-demos-server/services"
)

func main() {
	fmt.Println("Starting server")

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	data := services.GetAlbums()
	c.IndentedJSON(http.StatusOK, data)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	albums := services.GetAlbums()

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
