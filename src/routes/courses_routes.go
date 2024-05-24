package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/courses"
	"github.com/gin-gonic/gin"
)

func CoursesRoutes(g *gin.Engine, controller *courses.CourseController) {

	g.POST("/courses/create", controller.Create)
	g.GET("/courses", controller.GetAll)
	g.PUT("/courses/update", controller.UpdateCourse)
	g.GET("/courses/:id", controller.GetById)
}
