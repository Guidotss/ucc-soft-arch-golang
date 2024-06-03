package services

import (
	rating "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/rating"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/rating"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type IRatingService interface {
	NewRating(dto dto.RatingRequestResponseDto) (dto.RatingRequestResponseDto, error)
}

type ratingService struct {
	client rating.RatingClient
}

func NewRatingService(client *rating.RatingClient) IRatingService {
	return &ratingService{client: *client}
}

func (r *ratingService) NewRating(data dto.RatingRequestResponseDto) (dto.RatingRequestResponseDto, error) {
	var NewRating = model.Rating{
		CourseId: data.CourseId,
		UserId:   data.UserId,
		Rating:   data.Rating,
	}

	rating, err := r.client.NewRating(NewRating)
	if err != nil {
		return dto.RatingRequestResponseDto{}, err
	}

	return dto.RatingRequestResponseDto{
		CourseId: rating.CourseId,
		UserId:   rating.UserId,
		Rating:   rating.Rating,
	}, nil
}
