package inscriptos

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InscriptosClient struct {
	Db *gorm.DB
}

func NewInscriptionClient(db *gorm.DB) *InscriptosClient {
	return &InscriptosClient{Db: db}
}

func (c *InscriptosClient) Enroll(inscripto model.Inscriptos) model.Inscriptos {
	result := c.Db.Create(&inscripto)

	if result.Error != nil {
		log.Error()
	}
	return inscripto
}
