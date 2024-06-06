package rating

import (
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	model "github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"gorm.io/gorm"
)

type RatingClient struct {
	Db *gorm.DB
}

func NewRatingClient(db *gorm.DB) *RatingClient {
	return &RatingClient{Db: db}
}

func (c *RatingClient) NewRating(rating model.Rating) (model.Rating, error) {
	result := c.Db.Create(&rating)

	if result.Error != nil {
		return model.Rating{}, customError.NewError("INTERNAL_SERVER_ERROR", "Error creating rating", 500)
	}
	return rating, nil
}
