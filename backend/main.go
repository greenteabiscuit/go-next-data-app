package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/go-next-data-app/backend/controller"
)

func main() {
	r := gin.Default()
	helloController := controller.HelloController{}
	r.Use()
	{
		r.GET("/api/hello", helloController.Hello)
	}

	r.Run(":8080")
}
