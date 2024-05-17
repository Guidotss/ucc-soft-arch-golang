package courses

import (
	coursesDomain "github.com/Guidotss/ucc-soft-arch-golang.git/domain/dtos/courses"
	"github.com/Guidotss/ucc-soft-arch-golang.git/services"

	"github.com/gin-gonic/gin"
)

func Create(g *gin.Context) {
	var course coursesDomain.CreateCoursesRequestDto
	g.BindJSON(&course)
	var courseServices = services.NewCourseService()
	newCourse := courseServices.CreateCourse(course)

	g.JSON(201, newCourse)
}
