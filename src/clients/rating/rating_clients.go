package rating

import (
	model "github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RatingClient struct {
	Db *gorm.DB
}

func NewRatingClient(db *gorm.DB) *RatingClient {
	return &RatingClient{Db: db}
}

func (c *RatingClient) NewRating(rating model.Rating) model.Rating {
	result := c.Db.Create(&rating)

	if result.Error != nil {
		log.Error()
	}
	return rating
}

func (c *RatingClient) GetCourseRaiting(courseId uuid.UUID) model.Ratings {
	var ratings []model.Rating
	c.Db.Where("course_id = ?", courseId).Find(&ratings)
	var courseRatings model.Ratings
	for _, rating := range ratings {
		courseRatings = append(courseRatings, rating)
	}
	return courseRatings
}
