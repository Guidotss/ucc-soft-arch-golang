package adapter

import (
	client "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/courses"
	controllers "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/courses"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"gorm.io/gorm"
)

func CourseAdapter(db *gorm.DB) *controllers.CourseController {
	client := client.NewCourseClient(db)
	service := services.NewCourseService(client)
	return controllers.NewCourseController(service)
}
