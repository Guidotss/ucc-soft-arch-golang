package courses

import (
	"errors"
	"net/http"
	"strings"

	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type CourseClient struct {
	Db *gorm.DB
}

func NewCourseClient(db *gorm.DB) *CourseClient {
	return &CourseClient{Db: db}
}

func (c *CourseClient) Create(course model.Course) (model.Course, error) {
	result := c.Db.Create(&course)
	if result.Error != nil {
		var err error
		switch {
		case strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint"):
			err = customError.NewError(
				"DUPLICATE_IDENTIFIER",
				"A course with the same identifier already exists. Please use a different identifier.",
				http.StatusConflict)
		case strings.Contains(result.Error.Error(), "connection"):
			err = customError.NewError(
				"DB_CONNECTION_ERROR",
				"Database connection error. Please try again later.",
				http.StatusInternalServerError)
		default:
			err = customError.NewError(
				"UNEXPECTED_ERROR",
				"An unexpected error occurred. Please try again later.",
				http.StatusInternalServerError)
		}
		return model.Course{}, err
	}
	return course, nil
}

func (c *CourseClient) GetAll() (model.Courses, error) {
	var courses model.Courses
	var rawResults []map[string]interface{}
	err := c.Db.Raw(
		`SELECT courses.*, categories.category_name ,r.ratingavg
			FROM courses, 
				(SELECT course_id , AVG(rating) as ratingavg 
				 FROM ratings GROUP BY course_id) as r, 
				categories
			WHERE 
				courses.id = r.course_id AND 
				courses.category_id = categories.id`).Scan(&rawResults).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.NewError("NOT_FOUND", "There is no courses", http.StatusNotFound)
		}
		return nil, customError.NewError("DB_ERROR", "Error retrieving course from database", http.StatusInternalServerError)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Course{}, customError.NewError("NOT_FOUND", "Course not found", http.StatusNotFound)
		}
		return model.Course{}, customError.NewError("DB_ERROR", "Error retrieving course from database", http.StatusInternalServerError)
	}
	return course, nil
}
func (c *CourseClient) UpdateCourse(course model.Course) (model.Course, error) {

	result := c.Db.Table("courses").Where("id = ?", course.Id).Updates(&course)
	if result.Error != nil {
		var err error
		switch {
		case strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint"):
			err = customError.NewError(
				"DUPLICATE_IDENTIFIER",
				"A course with the same identifier or name already exists. Please use a different identifier or name.",
				http.StatusConflict)
		case strings.Contains(result.Error.Error(), "connection"):
			err = customError.NewError(
				"DB_CONNECTION_ERROR",
				"Database connection error. Please try again later.",
				http.StatusInternalServerError)
		default:
			err = customError.NewError(
				"UNEXPECTED_ERROR",
				"An unexpected error occurred. Please try again later.",
				http.StatusInternalServerError)
		}
		return model.Course{}, err
	}
	return course, nil
}

// FUNCION PARA PARSEAR UUID
func parseUUID(value interface{}) uuid.UUID {
	if value != nil {
		id, _ := uuid.Parse(value.(string))
		return id
	}
	return uuid.Nil
}
