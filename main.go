package main

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/config"
	"github.com/gin-gonic/gin"
)

func main() {
	envs := config.LoadEnvs(".env")
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + envs.Get("PORT"))
}
