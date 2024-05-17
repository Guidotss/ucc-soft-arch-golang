package courses

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Create(c model.Course) model.Course {
	result := Db.Create(&c)

	if result.Error != nil {
		log.Error()
	}
	log.Debug("Curso creado con exito wachin, su id es: ", result)
	return c
}
