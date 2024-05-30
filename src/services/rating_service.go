package services

import (
	rating "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/rating"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/rating"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type IRatingService interface {
	NewRating(dto dto.RatingRequestResponseDto) dto.RatingRequestResponseDto
	GetCourseRatings(courseId dto.GetCourseRatingRequestDto) dto.CourseRatingsDto
}

type ratingService struct {
	client rating.RatingClient
}

func NewRatingService(client *rating.RatingClient) IRatingService {
	return &ratingService{client: *client}
}

func (r *ratingService) NewRating(data dto.RatingRequestResponseDto) dto.RatingRequestResponseDto {
	var NewRating = model.Rating{
		CourseId: data.CourseId,
		UserId:   data.UserId,
		Rating:   data.Rating,
	}

	rating := r.client.NewRating(NewRating)

	return dto.RatingRequestResponseDto{
		CourseId: rating.CourseId,
		UserId:   rating.UserId,
		Rating:   rating.Rating,
	}
}
func (r *ratingService) GetCourseRatings(courseId dto.GetCourseRatingRequestDto) dto.CourseRatingsDto {
	var ratings = r.client.GetCourseRaiting(courseId.CourseId)
	var courseRatings dto.CourseRatingsDto
	for _, rating := range ratings {
		courseRatings = append(courseRatings, dto.GetCourseRatingResponseDto{
			CourseId: rating.CourseId,
			Rating:   rating.Rating,
		})
	}
	return courseRatings
}
