package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/rating"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/rating"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type IRatingService interface {
	NewRating(dto dto.RatingRequestResponseDto) dto.RatingRequestResponseDto
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
