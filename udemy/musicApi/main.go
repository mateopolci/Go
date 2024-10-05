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

// Controlador GET
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// Controlador GET un album por ID
func getAlbumById(ctx *gin.Context) {
	//Capturamos el valor pasado por endpoint
	id := ctx.Param("id")

	//Recorremos el array de albumes para ver si existe el id ingresado
	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	//Si no se encontró, devolvemos error
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message":"The album ID does not exist"})
}

// Controlador POST un álbum
func postAlbums(ctx *gin.Context) {
	var newAlbum album
	err := ctx.BindJSON(&newAlbum)
	
	if err != nil {
		//Devolver un error 400
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message":"The album ID does not exist"})
		return
	}
	//Añadimos el album al slice
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	//Establecemos el endpoint especificando el valor como id
	router.GET("/album/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
