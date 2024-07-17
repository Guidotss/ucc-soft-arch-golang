package rating

import "github.com/google/uuid"

type RatingRequestResponseDto struct {
	CourseId uuid.UUID `json:"course_id"`
	UserId   uuid.UUID `json:"user_id"`
	Rating   int       `json:"rating"`
}

type GetCourseRatingRequestDto struct {
	CourseId uuid.UUID `json:"course_id"`
}

type GetCourseRatingResponseDto struct {
	CourseId uuid.UUID `json:"course_id"`
	Rating   int       `json:"rating"`
}
type CourseRatingsDto []GetCourseRatingResponseDto
type RatingsResponse []RatingRequestResponseDto
