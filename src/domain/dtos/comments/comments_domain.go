package comments

import "github.com/google/uuid"

type CommentRequestResponseDto struct {
	CourseId uuid.UUID `json:"course_id"`
	UserId   uuid.UUID `json:"user_id"`
	Text     string    `json:"text"`
}

type GetCommentRequest struct {
	CourseId uuid.UUID `json:"course_id"`
}

type CommentResponse struct {
	Text string `json:"text"`
	User struct {
		Name   string    `json:"name"`
		Avatar string    `json:"avatar"`
		Id     uuid.UUID `json:"id"`
	}
}
type GetCommentsResponse []CommentResponse
