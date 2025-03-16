package main

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)

}

// postAlbum adds an album from JSON received in the request body
func postAlbum(c *gin.Context) {
	var newAlbum album

	//call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

// getAlbumByID locates the album whose ID value matches the id pararmeter sent by the client, then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, x := range albums {
		if x.ID == id {
			c.IndentedJSON(http.StatusOK, x)
			return
		}
	}
	//in case there is not that id in the slice "albums", an error message of 404 type will be shown
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for j, i := range albums {
		if i.ID == id {
			albums = slices.Delete(albums, j, j+1)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
