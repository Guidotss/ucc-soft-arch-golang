package model

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	cid               string  `gorm:"cid"`
	CourseName        string  `gorm:"course_name"`
	CourseDescription string  `gorm:"description"`
	CoursePrice       float64 `gorm:"price"`
	CourseDuration    int     `gorm:"duration"`
	CourseInitDate    string  `gorm:"init_date"`
	CourseState       bool    `gorm:"state"`
	CourseCapacity    int     `gorm:"cupo"`
	CourseImage       string  `gorm:"image"`
	CategoryID        string
	Category          Category
}
