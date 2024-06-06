package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/comments"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/comments"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
)

type ICommentsService interface {
	NewComment(dto dto.CommentRequestResponseDto) (dto.CommentRequestResponseDto, error)
	GetCourseComments(courseID uuid.UUID) (dto.GetCommentsResponse, error)
}

type commentsService struct {
	client comments.CommentsClient
}

func NewCommentsService(client *comments.CommentsClient) ICommentsService {
	return &commentsService{client: *client}
}

func (c *commentsService) NewComment(data dto.CommentRequestResponseDto) (dto.CommentRequestResponseDto, error) {
	var NewComment = model.Comment{
		CourseId: data.CourseId,
		UserId:   data.UserId,
		Text:     data.Text,
	}

	coment, err := c.client.NewComment(NewComment)
	if err != nil {
		return dto.CommentRequestResponseDto{}, err
	}

	return dto.CommentRequestResponseDto{
		CourseId: coment.CourseId,
		UserId:   coment.UserId,
		Text:     coment.Text,
	}, nil
}

func (c *commentsService) GetCourseComments(courseID uuid.UUID) (dto.GetCommentsResponse, error) {
	comments, err := c.client.GetCourseComments(courseID)
	if err != nil {
		return dto.GetCommentsResponse{}, err
	}
	var allCommentsDto dto.GetCommentsResponse
	for _, result := range comments {
		var commentDto dto.CommentResponse
		commentDto.User.Name = result.UserName
		commentDto.Text = result.Text
		commentDto.User.Avatar = result.UserAvatar
		commentDto.User.Id = result.UserId
		allCommentsDto = append(allCommentsDto, commentDto)
	}
	return allCommentsDto, nil

}
