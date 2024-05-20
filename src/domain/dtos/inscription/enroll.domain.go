package inscription

import "github.com/google/uuid"

type EnrollRequestResponseDto struct {
	CourseId uuid.UUID `json:"course_id"`
	UserId   uuid.UUID `json:"user_id"`
}
