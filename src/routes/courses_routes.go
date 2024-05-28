package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/courses"
	middlewareAdmin "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/admin"
	middlewareCourse "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/course"
	"github.com/gin-gonic/gin"
)

func CoursesRoutes(g *gin.Engine, controller *courses.CourseController) {

	g.POST("/courses/create", middlewareAdmin.AuthMiddleware(), controller.Create)
	g.GET("/courses", controller.GetAll)
	g.PUT("/courses/update", middlewareCourse.CheckCourseId(), middlewareAdmin.AuthMiddleware(), controller.UpdateCourse)
	g.GET("/courses/:id", controller.GetById)
}
