package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inscripto struct {
	gorm.Model
	CourseId uuid.UUID
	UserId   uuid.UUID

	User   User   `gorm:"foreignKey:UserId"`
	Course Course `gorm:"foreignKey:CourseId"`
}
