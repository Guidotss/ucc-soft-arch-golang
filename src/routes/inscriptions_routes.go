package routes

import (
	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/inscriptions"
	isAdmin "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/admin"
	enroll "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/enroll"
	isLogged "github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/user"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
)

func InscriptionsRoutes(g *gin.Engine, controller *controller.InscriptionController, service services.IInscriptionService) {

	g.POST("/enroll",
		isLogged.AuthMiddleware(),
		enroll.CourseExist(service),
		enroll.IsAlredyEnroll(service),
		controller.Create)

	g.GET("/myCourses/",
		isLogged.AuthMiddleware(),
		controller.GetMyCourses)

	g.GET("/studentsInThisCourse/:cid",
		isAdmin.AdminAuthMiddleware(),
		controller.GetMyStudents)

	g.GET("/isEnrolled/:cid",
		isLogged.AuthMiddleware(),
		controller.IsAlredyEnrolled)
}
