package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	CourseId uuid.UUID
	UserId   uuid.UUID
	Rating   int `gorm:"type:int"`

	User   User   `gorm:"foreignKey:UserId"`
	Course Course `gorm:"foreignKey:CourseId"`
}

type Ratings []Rating
