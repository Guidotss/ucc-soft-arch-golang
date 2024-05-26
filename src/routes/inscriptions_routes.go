package routes

import (
	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/inscriptions"
	enroll "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/enroll"
	isLogged "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/user"
	"github.com/gin-gonic/gin"
)

func InscriptionsRoutes(g *gin.Engine, controller *controller.InscriptionController) {

	g.POST("/enroll",
		isLogged.AuthMiddleware(),
		enroll.CourseExist(*controller),
		enroll.IsAlredyEnroll(*controller),
		controller.Create)

	g.GET("/myCourses/",
		isLogged.AuthMiddleware(),
		controller.GetMyCourses)

	g.GET("/studentsInThisCourse/:cid", controller.GetMyStudents)
}
