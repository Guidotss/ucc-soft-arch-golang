package comments

import (
	"errors"
	"net/http"
	"strings"

	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	model "github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentsClient struct {
	Db *gorm.DB
}

func NewCommentsClient(db *gorm.DB) *CommentsClient {
	return &CommentsClient{Db: db}
}

func (c *CommentsClient) NewComment(comment model.Comment) (model.Comment, error) {
	result := c.Db.Create(&comment)
	if result.Error != nil {
		var err error
		switch {
		case strings.Contains(result.Error.Error(), "connection"):
			err = customError.NewError(
				"DB_CONNECTION_ERROR",
				"Database connection error. Please try again later.",
				http.StatusInternalServerError)
		default:
			err = customError.NewError(
				"UNEXPECTED_ERROR",
				"An unexpected error occurred. Please try again later.",
				http.StatusInternalServerError)
		}
		return model.Comment{}, err
	}
	return comment, nil
}

func (c *CommentsClient) GetCourseComments(courseID uuid.UUID) (model.Comments, error) {
	var comments model.Comments
	var rawResults []map[string]interface{}
	err := c.Db.Raw(`
        SELECT C.Text, U.Name, U.Avatar, U.id as User_id
        FROM comments C
        JOIN users U ON C.user_id = U.id
        WHERE C.course_id = ?
    `, courseID).Scan(&rawResults).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.NewError("COMMENTS_NOT_FOUND", "No comments found for the specified course", http.StatusNotFound)
		} else if strings.Contains(err.Error(), "connection") {
			return nil, customError.NewError("DB_CONNECTION_ERROR", "Database connection error. Please try again later.", http.StatusInternalServerError)
		} else {
			return nil, customError.NewError("UNEXPECTED_ERROR", "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		}
	}
	if len(rawResults) == 0 {
		return nil, customError.NewError("COMMENTS_NOT_FOUND", "No comments found for the specified course", http.StatusNotFound)
	}

	for i := 0; i < len(rawResults); i++ {
		comment := model.Comment{
			Text:       rawResults[i]["text"].(string),
			UserName:   rawResults[i]["name"].(string),
			UserAvatar: rawResults[i]["avatar"].(string),
			UserId:     parseUUID(rawResults[i]["user_id"]),
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (c *CommentsClient) UpdateComment(comment model.Comment) (model.Comment, error) {
	result := c.Db.Table("comments").
		Where("user_id = ? AND course_id = ?", comment.UserId, comment.CourseId).
		Updates(&comment)
	if result.Error != nil {
		var err error
		switch {
		case strings.Contains(result.Error.Error(), "connection"):
			err = customError.NewError(
				"DB_CONNECTION_ERROR",
				"Database connection error. Please try again later.",
				http.StatusInternalServerError)
		default:
			err = customError.NewError(
				"UNEXPECTED_ERROR",
				"An unexpected error occurred. Please try again later.",
				http.StatusInternalServerError)
		}
		return model.Comment{}, err
	}
	return comment, nil
}

// FUNCION PARA PARSEAR UUID
func parseUUID(value interface{}) uuid.UUID {
	if value != nil {
		id, _ := uuid.Parse(value.(string))
		return id
	}
	return uuid.Nil
}
