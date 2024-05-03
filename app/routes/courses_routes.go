package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/controllers/courses"
	"github.com/gin-gonic/gin"
)

func CoursesRoutes(g *gin.RouterGroup /* controller *CoursesController */) {
	g.Group("/courses")
	{
		g.POST("/create", courses.Create)
	}
}
