package inscriptos

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InscriptosClient struct {
	Db *gorm.DB
}

func NewInscriptionClient(db *gorm.DB) *InscriptosClient {
	return &InscriptosClient{Db: db}
}

func (c *InscriptosClient) Enroll(inscripto model.Inscripto) model.Inscripto {
	result := c.Db.Create(&inscripto)

	if result.Error != nil {
		log.Error()
	}
	return inscripto
}

func (c *InscriptosClient) GetMyCourses(id uuid.UUID) model.MyCourses {
	var inscriptos []model.Inscripto
	c.Db.Where("user_id = ?", id).Find(&inscriptos)
	var courses model.MyCourses
	for _, inscripto := range inscriptos {
		courses = append(courses, inscripto.CourseId)
	}
	return courses
}
