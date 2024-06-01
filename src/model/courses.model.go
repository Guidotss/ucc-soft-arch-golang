package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Id                uuid.UUID `sql:"type:uuid;primary_key;default:gen_random_uuid()"`
	CourseName        string    `gorm:"course_name;unique"`
	CourseDescription string    `gorm:"description"`
	CoursePrice       float64   `gorm:"price"`
	CourseDuration    int       `gorm:"duration"`
	CourseInitDate    string    `gorm:"init_date"`
	CourseState       bool      `gorm:"state;default:false"`
	CourseCapacity    int       `gorm:"cupo;default:15"`
	CourseImage       string    `gorm:"image;default:https://upload.wikimedia.org/wikipedia/commons/a/a3/Image-not-found.png"`
	CategoryID        uuid.UUID
	Category          Category `gorm:"foreignKey:CategoryID"`
	Ratings           Ratings  `gorm:"foreignKey:CourseId"`
	RatingAvg         float64  `gorm:"-" json:"ratingavg"`
}

func (model *Course) BeforeCreate(tx *gorm.DB) (err error) {
	model.Id = uuid.New()
	return
}

type Courses []Course
