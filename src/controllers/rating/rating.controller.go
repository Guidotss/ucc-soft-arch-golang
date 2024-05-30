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
	if err := g.ShouldBindJSON(&ratingDto); err != nil {
		g.JSON(400, gin.H{
			"message": "La estructura del json es incorrecta",
		})
		return
	}

	response := c.RatingService.NewRating(ratingDto)
	g.JSON(201, gin.H{
		"response": response,
		"message":  "La valoracion se registro con exito",
	})
}
func (c *RatingController) GetCourseRaiting(g *gin.Context) {
	var courseId dto.GetCourseRatingRequestDto
	if err := g.BindJSON(&courseId); err != nil {
		g.JSON(400, gin.H{
			"message": "La estructura del json es incorrecta",
		})
		return
	}
	response := c.RatingService.GetCourseRatings(courseId)
	RatingQty := len(response)
	if RatingQty == 0 {
		g.JSON(404, gin.H{
			"message": "No se encontraron valoraciones para el curso",
		})
		return
	}
	var RatingTotal int
	for _, rating := range response {
		RatingTotal += rating.Rating
	}
	RatingAvg := float64(RatingTotal) / float64(RatingQty)

	g.JSON(200, gin.H{
		"response": response,
		"quantity": RatingQty,
		"avg":      RatingAvg,
	})
}
