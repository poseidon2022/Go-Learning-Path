package main


import (
	"net/http"
	"github.com/gin-gonic/gin"
)


//this section is all about the gin framework and backend apis

//build a simple end point application that handles music addition, deletion and further more
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, albums)
	}
}

func postAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var addedAlbum album
		if err := c.ShouldBind(&addedAlbum); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"fill out all the required fields"})
			return
		}

		wanted := addedAlbum.ID
		for _, value := range albums {
			if value.ID == wanted {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"ID duplicate found"})
				return
			}
		}

		albums = append(albums, addedAlbum)
		c.IndentedJSON(http.StatusOK, albums)
	}
}

func findByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		album_id := c.Param("id")
		for _, value := range albums {
			if value.ID == album_id {
				c.IndentedJSON(http.StatusOK, value)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})
	}
}

func main() {
	router := gin.Default()
	router.GET("/all-albums", getAll())
	router.POST("/new-album", postAlbum())

	//getting a specific item with an ID
	router.GET("/albums/:id", findByID())

	router.Run("localhost:8080")
}