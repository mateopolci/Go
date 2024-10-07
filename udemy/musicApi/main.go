package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{
		ID:     "1",
		Title:  "Abbey Road",
		Artist: "The Beatles",
		Price:  12.99,
	},
	{
		ID:     "2",
		Title:  "The Dark Side of the Moon",
		Artist: "Pink Floyd",
		Price:  14.99,
	},
	{
		ID:     "3",
		Title:  "Rumours",
		Artist: "Fleetwood Mac",
		Price:  11.99,
	},
}

// Controlador GET todo
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// Controlador GET un album por ID
func getAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message":"The album ID does not exist"})
}

// Controlador POST un álbum
func postAlbums(ctx *gin.Context) {
	var newAlbum album
	err := ctx.BindJSON(&newAlbum)
	
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message":"The album ID does not exist"})
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
