package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	CourseId uuid.UUID
	UserId   uuid.UUID
	Text     string `gorm:"text"`

	User   User
	Course Course
}
