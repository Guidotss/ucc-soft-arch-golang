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

	// Obtener userID y courseID de gin.Context y manejarlos correctamente
	userID, exists := g.Get("userID")
	if !exists {
		g.JSON(http.StatusBadRequest, gin.H{"error": "userID is required"})
		return
	}

	courseID, exists := g.Get("courseID")
	if !exists {
		g.JSON(http.StatusBadRequest, gin.H{"error": "courseID is required"})
		return
	}

	// Verificar y convertir userID a uuid.UUID
	var uid uuid.UUID
	switch v := userID.(type) {
	case string:
		parsedUID, err := uuid.Parse(v)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid userID format"})
			return
		}
		uid = parsedUID
	case []byte:
		parsedUID, err := uuid.ParseBytes(v)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid userID format"})
			return
		}
		uid = parsedUID
	case uuid.UUID:
		uid = v
	default:
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid userID format"})
		return
	}

	// Verificar y convertir courseID a uuid.UUID
	var cid uuid.UUID
	switch v := courseID.(type) {
	case string:
		parsedCID, err := uuid.Parse(v)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid courseID format"})
			return
		}
		cid = parsedCID
	case []byte:
		parsedCID, err := uuid.ParseBytes(v)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "invalid courseID format"})
			return
		}
		cid = parsedCID
	case uuid.UUID:
		cid = v
	default:
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid courseID format"})
		return
	}

	// Asignar valores a enrollDto
	enrollDto.UserId = uid
	enrollDto.CourseId = cid

	// Llamar al servicio de inscripción
	response, err := c.InscriptionService.Enroll(enrollDto)
	if err != nil {
		g.Error(err)
		return
	}

	// Responder con éxito
	g.JSON(http.StatusCreated, gin.H{
		"response": response,
		"message":  "El usuario se registró con éxito",
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

func (c *InscriptionController) IsAlredyEnrolled(g *gin.Context) {
	cid := g.Param("cid")
	course_id := parseUUID(cid)
	uid, _ := g.Get("userID")
	user_id := parseUUID(uid)
	flag, _ := c.InscriptionService.IsUserEnrolled(user_id, course_id)
	if flag {
		g.Error(customError.NewError("USER_ALREADY_ENROLLED", "User is already enrolled", http.StatusBadRequest))
	}
	g.JSON(200, gin.H{"message": "User is not enrolled"})
}
func (c *InscriptionController) CourseExist(course_id uuid.UUID) bool {
	flag, _ := c.InscriptionService.CourseExist(course_id)
	return flag
}

// FUNCION PARA PARSEAR UUID
func parseUUID(value interface{}) uuid.UUID {
	if value != nil {
		id, _ := uuid.Parse(value.(string))
		return id
	}
	return uuid.Nil
}
