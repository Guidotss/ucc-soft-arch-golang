package comments

import "github.com/google/uuid"

type CommentRequestResponseDto struct {
	CourseId uuid.UUID `json:"course_id"`
	UserId   uuid.UUID `json:"user_id"`
	Text     string    `json:"text"`
}
