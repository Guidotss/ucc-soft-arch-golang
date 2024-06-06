package inscriptions

import (
	"net/http"

	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/inscription"
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"

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
	userID, _ := g.Get("userID")
	courseID, _ := g.Get("courseID")
	uid := userID.(uuid.UUID)
	cid := courseID.(uuid.UUID)

	enrollDto.UserId = uid
	enrollDto.CourseId = cid

	response, err := c.InscriptionService.Enroll(enrollDto)
	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(201, gin.H{
		"response": response,
		"message":  "El usuarios se registro con exito",
	})
}

func (c *InscriptionController) GetMyCourses(g *gin.Context) {
	userID, _ := g.Get("userID")
	id := userID.(uuid.UUID)

	response, err := c.InscriptionService.GetMyCourses(id)
	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(200, response)
}
func (c *InscriptionController) GetMyStudents(g *gin.Context) {
	id := g.Param("cid")
	uuid, err := uuid.Parse(id)
	if err != nil {
		g.Error(customError.NewError("INVALID_UUID", "Invalid UUID", http.StatusBadRequest))
		return
	}

	response, err := c.InscriptionService.GetMyStudents(uuid)
	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(200, response)
}

// MIDDLEWARE FUNC
func (c *InscriptionController) IsAlredyEnrolled(user_id uuid.UUID, course_id uuid.UUID) bool {
	flag, _ := c.InscriptionService.IsUserEnrolled(user_id, course_id)
	return flag
}
func (c *InscriptionController) CourseExist(course_id uuid.UUID) bool {
	flag, _ := c.InscriptionService.CourseExist(course_id)
	return flag
}
