package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/courses"
	middleware "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/admin"
	"github.com/gin-gonic/gin"
)

func CoursesRoutes(g *gin.Engine, controller *courses.CourseController) {

	g.POST("/courses/create", middleware.AuthMiddleware(), controller.Create)
	g.GET("/courses", controller.GetAll)
	g.PUT("/courses/update", middleware.AuthMiddleware(), controller.UpdateCourse)
	g.GET("/courses/:id", controller.GetById)
}
