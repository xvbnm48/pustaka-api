package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "sakura endo",
		"bio":  "golang developer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content":  "teknologi",
		"subtitle": "golang",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
