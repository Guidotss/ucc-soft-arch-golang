package routes

import (
	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/inscriptions"
	"github.com/gin-gonic/gin"
)

func InscriptionsRoutes(g *gin.Engine, controller *controller.InscriptionController) {

	g.POST("/enroll", controller.Create)
	g.GET("/myCourses/:id", controller.GetMyCourses)
}
