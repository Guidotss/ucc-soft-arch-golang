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
	Text        string    `json:"comment"`
	User_name   string    `json:"user_name"`
	User_avatar string    `json:"user_avatar"`
	User_id     uuid.UUID `json:"user_id"`
}
type GetCommentsResponse []CommentResponse
