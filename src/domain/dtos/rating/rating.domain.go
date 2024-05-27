package rating

import "github.com/google/uuid"

type RatingRequestResponseDto struct {
	CourseId uuid.UUID `json:"course_id"`
	UserId   uuid.UUID `json:"user_id"`
	Rating   int       `json:"rating"`
}
