package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Running server!")
	server := gin.New()
	server.Run(":8080")
}
