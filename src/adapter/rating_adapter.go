package adapter

import (
	client "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/rating"
	controllers "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/rating"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"gorm.io/gorm"
)

func RatingAdapter(db *gorm.DB) *controllers.RatingController {
	client := client.NewRatingClient(db)
	service := services.NewRatingService(client)
	return controllers.NewRatingController(service)
}
