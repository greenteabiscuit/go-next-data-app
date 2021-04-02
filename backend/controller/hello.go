package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// HelloController ...
type HelloController struct {
}

// Download ...
func (c *HelloController) Hello(ctx *gin.Context) {
	fmt.Println("hello")
	return
}
