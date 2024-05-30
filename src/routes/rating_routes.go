package routes

import (
	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/rating"
	"github.com/gin-gonic/gin"
)

func RatingRoutes(g *gin.Engine, controller *controller.RatingController) {
	g.POST("/rating", controller.NewRating)
	g.GET("/courseRatings", controller.GetCourseRaiting)
}
