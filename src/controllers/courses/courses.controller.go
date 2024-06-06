package courses

import (
	"net/http"

	coursesDomain "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/courses"
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
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
	if err := g.BindJSON(&courseDto); err != nil {
		g.Error(err)
		return
	}
	response, err := c.CourseService.CreateCourse(courseDto)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(201, gin.H{
		"ok":      true,
		"message": "Course created successfully",
		"data":    response,
	})
}

func (c *CourseController) GetAll(g *gin.Context) {
	response, err := c.CourseService.FindAllCourses()
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(200, response)
}

func (c *CourseController) GetById(g *gin.Context) {
	id := g.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		g.Error(customError.NewError("INVALID_UUID", "Invalid UUID", http.StatusBadRequest))
		return
	}

	response, err := c.CourseService.FindOneCourse(uuid)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(200, response)
}

func (c *CourseController) UpdateCourse(g *gin.Context) {
	var courseDto coursesDomain.UpdateRequestDto
	if err := g.BindJSON(&courseDto); err != nil {
		g.Error(err)
		return
	}
	response, err := c.CourseService.UpdateCourse(courseDto)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(201, gin.H{
		"ok":      true,
		"message": "Course updated successfully",
		"data":    response,
	})
}
