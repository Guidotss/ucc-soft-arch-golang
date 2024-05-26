package inscriptions

import (
	"fmt"
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
	userID, exists := g.Get("userID")
	if !exists {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}
	courseID, exists := g.Get("courseID")
	if !exists {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Course ID not found"})
		return
	}
	uid := userID.(uuid.UUID)
	cid, err := uuid.Parse(courseID.(string))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Course ID"})
		return
	}
	enrollDto.UserId = uid
	enrollDto.CourseId = cid

	response := c.InscriptionService.Enroll(enrollDto)
	g.JSON(201, gin.H{
		"response": response,
		"message":  "El usuarios se registro con exito",
	})
}

func (c *InscriptionController) GetMyCourses(g *gin.Context) {
	userID, exists := g.Get("userID")
	if !exists {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}
	id := userID.(uuid.UUID)
	response := c.InscriptionService.GetMyCourses(id)
	g.JSON(200, response)
}
func (c *InscriptionController) GetMyStudents(g *gin.Context) {
	id := g.Param("cid")
	uuid, err := uuid.Parse(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	response := c.InscriptionService.GetMyStudents(uuid)
	g.JSON(200, response)
}

// MIDDLEWARE FUNC
func (c *InscriptionController) IsAlredyEnrolled(user_id uuid.UUID, course_id uuid.UUID) bool {
	flag, _ := c.InscriptionService.IsUserEnrolled(user_id, course_id)
	fmt.Println("controller enrolled flag: ", flag)
	return flag
}
func (c *InscriptionController) CourseExist(course_id uuid.UUID) bool {
	flag, _ := c.InscriptionService.CourseExist(course_id)
	fmt.Println("controller existcourse flag: ", flag)
	return flag
}
