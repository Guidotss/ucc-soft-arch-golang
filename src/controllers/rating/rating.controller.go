package rating

import (
	"net/http"

	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/rating"
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
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
	if err := g.ShouldBindJSON(&ratingDto); err != nil {
		g.Error(customError.NewError("INVALID_INPUTS", "Invalid fields", http.StatusBadRequest))
		return
	}

	response, err := c.RatingService.NewRating(ratingDto)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(201, gin.H{
		"response": response,
		"message":  "La valoracion se registro con exito",
	})
}
func (c *RatingController) UpdateRating(g *gin.Context) {
	var ratingDto dto.RatingRequestResponseDto
	if err := g.ShouldBindJSON(&ratingDto); err != nil {
		g.Error(customError.NewError("INVALID_INPUTS", "Invalid fields", http.StatusBadRequest))
		return
	}
	response, err := c.RatingService.UpdateRating(ratingDto)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(200, gin.H{
		"response": response,
		"message":  "La valoracion se actualizo con exito",
	})
}
func (c *RatingController) GetRatings(g *gin.Context) {
	response, err := c.RatingService.GetRatings()
	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(200, response)
}
