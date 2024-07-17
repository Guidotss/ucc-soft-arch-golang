package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text       string `gorm:"type:text"`
	UserId     uuid.UUID
	CourseId   uuid.UUID
	User       User   `gorm:"foreignKey:UserId"`
	Course     Course `gorm:"foreignKey:CourseId"`
	UserName   string `gorm:"-"`
	UserAvatar string `gorm:"-"`
}

type Comments []Comment
