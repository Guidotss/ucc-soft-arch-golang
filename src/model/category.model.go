package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id           uuid.UUID `sql:"type:uuid;primary_key;default:gen_random_uuid()"`
	CategoryName string    `gorm:"category_name"`
}

func (model *Category) BeforeCreate(tx *gorm.DB) (err error) {
	model.Id = uuid.New()
	return
}

type Categories []Category
