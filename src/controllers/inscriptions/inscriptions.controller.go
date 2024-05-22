package inscriptions

import (
	"net/http"

	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/inscription"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (c *InscriptionController) GetMyCourses(g *gin.Context) {
	id := g.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	response := c.InscriptionService.GetMyCourses(uuid)
	g.JSON(200, response)

}
