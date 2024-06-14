package inscriptos

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InscriptosClient struct {
	Db *gorm.DB
}

func NewInscriptionClient(db *gorm.DB) *InscriptosClient {
	return &InscriptosClient{Db: db}
}

func (c *InscriptosClient) Enroll(inscripto model.Inscripto) (model.Inscripto, error) {
	result := c.Db.Create(&inscripto)

	if result.Error != nil {
		var err error
		switch {
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
		return model.Inscripto{}, err
	}
	return inscripto, nil
}

func (c *InscriptosClient) GetMyCourses(id uuid.UUID) (model.Courses, error) {
	var rawResults []map[string]interface{}
	err := c.Db.Raw(`
	SELECT C.*, CAT.category_name
		FROM courses C
		JOIN inscriptos I ON I.course_id = C.id
		JOIN users U ON I.user_id = U.id
		JOIN categories CAT ON C.category_id = CAT.id
		WHERE I.user_id = ?
	`, id).Scan(&rawResults).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.NewError("COMMENTS_NOT_FOUND", "No Courses found for the specified user", http.StatusNotFound)
		} else if strings.Contains(err.Error(), "connection") {
			return nil, customError.NewError("DB_CONNECTION_ERROR", "Database connection error. Please try again later.", http.StatusInternalServerError)
		} else {
			return nil, customError.NewError("UNEXPECTED_ERROR", "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		}
	}
	var courses model.Courses
	for i := 0; i < len(rawResults); i++ {
		course := model.Course{
			Id:                parseUUID(rawResults[i]["id"]),
			CourseName:        rawResults[i]["course_name"].(string),
			CourseDescription: rawResults[i]["course_description"].(string),
			CoursePrice:       rawResults[i]["course_price"].(float64),
			CourseDuration:    int(rawResults[i]["course_duration"].(int64)),
			CourseInitDate:    rawResults[i]["course_init_date"].(string),
			CourseState:       rawResults[i]["course_state"].(bool),
			CourseCapacity:    int(rawResults[i]["course_capacity"].(int64)),
			CourseImage:       rawResults[i]["course_image"].(string),
			CategoryID:        parseUUID(rawResults[i]["category_id"]),
			Category: model.Category{
				CategoryName: rawResults[i]["category_name"].(string),
			},
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *InscriptosClient) GetMyStudents(id uuid.UUID) (model.Users, error) {
	var rawResults []map[string]interface{}
	err := c.Db.Raw(`
		SELECT  U.name, U.avatar, U.id as User_id
		FROM inscriptos I, users U
		WHERE I.user_id = U.id AND I.course_id = ?
	`, id).Scan(&rawResults).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.NewError("STUDENST_NOT_FOUND", "No Students found for the specified course", http.StatusNoContent)
		} else if strings.Contains(err.Error(), "connection") {
			return nil, customError.NewError("DB_CONNECTION_ERROR", "Database connection error. Please try again later.", http.StatusInternalServerError)
		} else {
			return nil, customError.NewError("UNEXPECTED_ERROR", "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		}
	}
	var students model.Users
	for i := 0; i < len(rawResults); i++ {
		student := model.User{
			Name:   rawResults[i]["name"].(string),
			Avatar: rawResults[i]["avatar"].(string),
			Id:     parseUUID(rawResults[i]["user_id"]),
		}
		students = append(students, student)
	}
	return students, nil
}

// MIDDLEWARE FUNC
func (c *InscriptosClient) IsUserEnrolled(userID uuid.UUID, courseID uuid.UUID) (bool, error) {
	var count int64
	err := c.Db.Model(&model.Inscripto{}).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	fmt.Println("Enrolled count: ", count)
	return count > 0, nil
}
func (c *InscriptosClient) CourseExist(course_id uuid.UUID) (bool, error) {
	var count int64
	err := c.Db.Model(&model.Course{}).
		Where("Id = ?", course_id).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	fmt.Println("Count: ", count)
	return count > 0, nil
}

// FUNCION PARA PARSEAR UUID
func parseUUID(value interface{}) uuid.UUID {
	if value != nil {
		id, _ := uuid.Parse(value.(string))
		return id
	}
	return uuid.Nil
}
