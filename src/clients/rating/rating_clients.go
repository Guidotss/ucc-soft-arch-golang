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
func (c *RatingClient) UpdateRating(rating model.Rating) (model.Rating, error) {
	result := c.Db.Table("ratings").
		Where("user_id = ? AND course_id = ?", rating.UserId, rating.CourseId).
		Updates(&rating)
	if result.Error != nil {
		return model.Rating{}, customError.NewError("INTERNAL_SERVER_ERROR", "Error updating rating", 500)
	}
	return rating, nil
}
func (c *RatingClient) GetRatings() (model.Ratings, error) {
	var ratings model.Ratings
	result := c.Db.Find(&ratings)
	if result.Error != nil {
		return nil, customError.NewError("INTERNAL_SERVER_ERROR", "Error getting ratings", 500)
	}
	return ratings, nil
}
