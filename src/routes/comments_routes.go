package routes

import (
	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/comments"
	"github.com/gin-gonic/gin"
)

func CommentsRoutes(g *gin.Engine, controller *controller.CommentsController) {
	g.POST("/comment", controller.NewComment)
	g.GET("/comment/:id", controller.GetCourseComments)
	g.PUT("/comment", controller.UpdateComment)
}
