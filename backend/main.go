package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	x := 1
	y := "a"
	fmt.Println(x, y)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	r.Run(":8080")
}
