package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/courses"
	middlewareAdmin "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/admin"
	middlewareCourse "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/course"
	"github.com/gin-gonic/gin"
)

func CoursesRoutes(g *gin.Engine, controller *courses.CourseController) {

	g.POST("/courses/create",
		middlewareAdmin.AdminAuthMiddleware(),
		controller.Create)
	g.GET("/courses", controller.GetAll)
	g.PUT("/courses/update/:id",
		middlewareCourse.CheckCourseId(),
		middlewareAdmin.AdminAuthMiddleware(),
		controller.UpdateCourse)
	g.GET("/courses/:id", controller.GetById)
	g.DELETE("/courses/:id",
		middlewareAdmin.AdminAuthMiddleware(),
		controller.DeleteCourse)
}
