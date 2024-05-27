package rating

import (
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/rating"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
)

type RatingController struct {
	RatingService services.IRatingService
}

func NewRatingController(service services.IRatingService) *RatingController {
	return &RatingController{RatingService: service}
}

func (c *RatingController) NewRating(g *gin.Context) {
	var ratingDto dto.RatingRequestResponseDto

	response := c.RatingService.NewRating(ratingDto)
	g.JSON(201, gin.H{
		"response": response,
		"message":  "La valoracion se registro con exito",
	})
}
