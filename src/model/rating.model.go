package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	CourseId uuid.UUID
	UserId   uuid.UUID
	Rating   int `gorm:"rating"`

	User   User
	Course Course
}
