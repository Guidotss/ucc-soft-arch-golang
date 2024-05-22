package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inscripto struct {
	gorm.Model
	CourseId uuid.UUID
	UserId   uuid.UUID

	User   User
	Course Course
}

type StudentsInCourse []uuid.UUID
type MyCourses []uuid.UUID
