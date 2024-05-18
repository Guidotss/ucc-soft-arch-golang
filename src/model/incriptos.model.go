package model

import (
	"gorm.io/gorm"
)

type Inscriptos struct {
	gorm.Model
	CourseId string
	UserId   string

	User   User
	Course Course
}
