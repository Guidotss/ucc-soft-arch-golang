package inscriptions

import (
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/inscription"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
)

type InscriptionController struct {
	InscriptionService services.IInscriptionService
}

func NewInscriptionController(service services.IInscriptionService) *InscriptionController {
	return &InscriptionController{InscriptionService: service}
}

func (c *InscriptionController) Create(g *gin.Context) {
	var enrollDto dto.EnrollRequestResponseDto
	err := g.BindJSON(&enrollDto)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.InscriptionService.Enroll(enrollDto)
	g.JSON(201, gin.H{
		"response": response,
		"message":  "El usuarios se registro con exito",
	})
}
