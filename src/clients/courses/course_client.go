package courses

import (
	"fmt"

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

func (c *CourseClient) GetAll() (model.Courses, error) {
	var courses model.Courses
	var rawResults []map[string]interface{}
	err := c.Db.Raw(
		`
			SELECT courses.*, categories.category_name ,r.ratingavg
			FROM courses, 
				(SELECT course_id , AVG(rating) as ratingavg 
				 FROM ratings GROUP BY course_id) as r, 
				categories
			WHERE courses.id = r.course_id AND courses.category_id = categories.id		
		`).Scan(&rawResults).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, data := range rawResults {
		course := model.Course{
			Id:                parseUUID(data["id"]),
			CourseName:        data["course_name"].(string),
			CourseDescription: data["course_description"].(string),
			CoursePrice:       data["course_price"].(float64),
			CourseDuration:    int(data["course_duration"].(int64)),
			CourseInitDate:    data["course_init_date"].(string),
			CourseState:       data["course_state"].(bool),
			CourseCapacity:    int(data["course_capacity"].(int64)),
			CourseImage:       data["course_image"].(string),
			CategoryID:        parseUUID(data["category_id"]),
			Category: model.Category{
				CategoryName: data["category_name"].(string),
			},
			RatingAvg: data["ratingavg"].(float64),
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *CourseClient) GetById(id uuid.UUID) (model.Course, error) {
	var course model.Course
	err := c.Db.Where("id = ?", id).First(&course).Error
	if err != nil {
		return model.Course{}, err
	}
	return course, nil
}
func (c *CourseClient) UpdateCourse(dto dto.UpdateRequestDto) (model.Course, error) {
	var course model.Course
	result := c.Db.First(&course, dto.Id)

	if result.Error != nil {
		return model.Course{}, result.Error
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
	if dto.CourseImage != nil {
		course.CourseImage = *dto.CourseImage
	}

	result = c.Db.Save(&course)
	if result.Error != nil {
		return model.Course{}, result.Error
	}

	return course, nil
}

// FUNCIONES PARA PARSEAR TIPOS
func parseUUID(value interface{}) uuid.UUID {
	if value != nil {
		id, _ := uuid.Parse(value.(string))
		return id
	}
	return uuid.Nil
}

/*
func parseString(value interface{}) string {
	if value != nil {
		return value.(string)
	}
	return ""
}

func parseFloat(value interface{}) float64 {
	if value != nil {
		return value.(float64)
	}
	return 0.0
}

func parseInt(value interface{}) int {
	if value != nil {
		return int(value.(int64))
	}
	return 0
}

func parseBool(value interface{}) bool {
	if value != nil {
		return value.(bool)
	}
	return false
}
*/
