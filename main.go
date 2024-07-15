package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func showGetEndpoint(c *gin.Context) {

	showData := gin.H{
		"message": "pong",
	}

	c.JSON(200, showData)
}

func main() {
	fmt.Println("Hello, Go!")

	r := gin.Default()
	r.GET("/ping", showGetEndpoint)
	r.Run(":5000")
}
