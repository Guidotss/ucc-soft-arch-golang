package courses

import (
	coursesDomain "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/courses"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	CourseService services.ICourseService
}

type ICourseController interface {
	Create(g *gin.Context)
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
