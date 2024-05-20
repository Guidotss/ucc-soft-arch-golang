package courses

import (
	"net/http"

	coursesDomain "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/courses"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CourseController struct {
	CourseService services.ICourseService
}

func NewCourseController(service services.ICourseService) *CourseController {
	return &CourseController{CourseService: service}
}

func (c *CourseController) Create(g *gin.Context) {
	var courseDto coursesDomain.CreateCoursesRequestDto
	err := g.BindJSON(&courseDto)
	if err != nil {
		g.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := c.CourseService.CreateCourse(courseDto)
	g.JSON(201, response)
}

func (c *CourseController) GetAll(g *gin.Context) {
	response := c.CourseService.FindAllCourses()
	g.JSON(200, response)
}

func (c *CourseController) GetById(g *gin.Context) {
	id := g.Param("id")
	uuid, err := uuid.Parse(id)
	response := c.CourseService.FindOneCourse(uuid)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	g.JSON(200, response)
}
