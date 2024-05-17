package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/controllers/courses"
	"github.com/gin-gonic/gin"
)

func CoursesRoutes(g *gin.Engine /* controller *CoursesController */) {

	g.POST("/courses/create", courses.Create)

}
