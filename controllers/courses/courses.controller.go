package courses

import (
	coursesDomain "github.com/Guidotss/ucc-soft-arch-golang.git/domain/dtos/courses"
	coursesServices "github.com/Guidotss/ucc-soft-arch-golang.git/services"

	"github.com/gin-gonic/gin"
)

func Create(g *gin.Context) {
	var course coursesDomain.CoursesCreateRequest
	g.BindJSON(&course)
	coursesServices.Create(course)
	g.JSON(201, gin.H{
		"message": "Create",
	})
}
