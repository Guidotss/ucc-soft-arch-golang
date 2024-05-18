package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inscriptos struct {
	gorm.Model
	CourseId uuid.UUID
	UserId   uuid.UUID

	User   User
	Course Course
}
