package api

import "github.com/gin-gonic/gin"

func ApiRoutes(g *gin.Engine) {
	routes := g.Group("/api")

	routes.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Test API",
		})
	})

}
