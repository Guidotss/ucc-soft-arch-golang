package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/courses"
	"github.com/gin-gonic/gin"
)

func CoursesRoutes(g *gin.Engine, controller *courses.CourseController) {

	g.POST("/courses/create", controller.Create)

}
