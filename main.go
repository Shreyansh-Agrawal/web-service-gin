package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type album struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}


var albums = []album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}


func getAlbumsById(context *gin.Context) {
	id := context.Param("id")

	for _, album := range albums {
		if album.Id == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "album not found"})
}


func postAlbums(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusOK, newAlbum)
}


func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
