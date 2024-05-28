package comments

import (
	model "github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CommentsClient struct {
	Db *gorm.DB
}

func NewCommentsClient(db *gorm.DB) *CommentsClient {
	return &CommentsClient{Db: db}
}

func (c *CommentsClient) NewComment(comments model.Comment) model.Comment {
	result := c.Db.Create(&comments)

	if result.Error != nil {
		log.Error()
	}
	return comments
}
