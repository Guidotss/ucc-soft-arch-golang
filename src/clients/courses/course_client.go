package courses

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CourseClient struct {
	Db *gorm.DB
}

func NewCourseClient(db *gorm.DB) *CourseClient {
	return &CourseClient{Db: db}
}

func (c *CourseClient) Create(course model.Course) model.Course {
	result := c.Db.Create(&course)

	if result.Error != nil {
		log.Error()
	}
	log.Debug("Curso creado con exito wachin, su id es: ", result)
	return course
}

func (c *CourseClient) GetAll() model.Courses {
	var courses model.Courses
	c.Db.Find(&courses)
	return courses
}

func (c *CourseClient) GetById(id uuid.UUID) model.Course {
	var course model.Course
	c.Db.Where("id = ?", id).First(&course)
	return course
}
