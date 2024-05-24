package courses

import (
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/courses"
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
func (c *CourseClient) UpdateCourse(dto dto.UpdateRequestDto) model.Course {
	var course model.Course
	result := c.Db.First(&course, dto.Id)

	if result.Error != nil {
		//manejo de errores
		panic(result.Error)
	}
	if dto.CourseName != nil {
		course.CourseName = *dto.CourseName
	}
	if dto.CourseDescription != nil {
		course.CourseDescription = *dto.CourseDescription
	}
	if dto.CoursePrice != nil {
		course.CoursePrice = *dto.CoursePrice
	}
	if dto.CourseDuration != nil {
		course.CourseDuration = *dto.CourseDuration
	}
	if dto.CourseCapacity != nil {
		course.CourseCapacity = *dto.CourseCapacity
	}
	if dto.CategoryID != nil {
		course.CategoryID = *dto.CategoryID
	}
	if dto.CourseInitDate != nil {
		course.CourseInitDate = *dto.CourseInitDate
	}
	if dto.CourseState != nil {
		course.CourseState = *dto.CourseState
	}
	result = c.Db.Save(&course)
	if result.Error != nil {
		//manejo de errores
		panic(result.Error)
	}
	return course
}
