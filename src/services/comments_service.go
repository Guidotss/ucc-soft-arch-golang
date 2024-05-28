package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/comments"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/comments"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type ICommentsService interface {
	NewComment(dto dto.CommentRequestResponseDto) dto.CommentRequestResponseDto
}

type commentsService struct {
	client comments.CommentsClient
}

func NewCommentsService(client *comments.CommentsClient) ICommentsService {
	return &commentsService{client: *client}
}

func (r *commentsService) NewComment(data dto.CommentRequestResponseDto) dto.CommentRequestResponseDto {
	var NewComment = model.Comment{
		CourseId: data.CourseId,
		UserId:   data.UserId,
		Text:     data.Text,
	}

	coment := r.client.NewComment(NewComment)

	return dto.CommentRequestResponseDto{
		CourseId: coment.CourseId,
		UserId:   coment.UserId,
		Text:     coment.Text,
	}
}
