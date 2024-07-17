package routes

import (
	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/rating"
	"github.com/gin-gonic/gin"
)

func RatingRoutes(g *gin.Engine, controller *controller.RatingController) {
	g.POST("/rating", controller.NewRating)
	g.PUT("/rating", controller.UpdateRating)
	g.GET("/rating", controller.GetRatings)
}
