package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Id                uuid.UUID `sql:"type:uuid;primary_key;default:gen_random_uuid()"`
	CourseName        string    `gorm:"course_name"`
	CourseDescription string    `gorm:"description"`
	CoursePrice       float64   `gorm:"price"`
	CourseDuration    int       `gorm:"duration"`
	CourseInitDate    string    `gorm:"init_date"`
	CourseState       bool      `gorm:"state"`
	CourseCapacity    int       `gorm:"cupo"`
	CourseImage       string    `gorm:"image"`
	CategoryID        uuid.UUID
	Category          Category
}

func (model *Course) BeforeCreate(tx *gorm.DB) (err error) {
	model.Id = uuid.New()
	return
}
