package inscriptos

import (
	"fmt"

	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InscriptosClient struct {
	Db *gorm.DB
}

func NewInscriptionClient(db *gorm.DB) *InscriptosClient {
	return &InscriptosClient{Db: db}
}

func (c *InscriptosClient) Enroll(inscripto model.Inscripto) model.Inscripto {
	result := c.Db.Create(&inscripto)

	if result.Error != nil {
		log.Error()
	}
	return inscripto
}

func (c *InscriptosClient) GetMyCourses(id uuid.UUID) model.MyCourses {
	var inscriptos []model.Inscripto
	c.Db.Where("user_id = ?", id).Find(&inscriptos)
	var courses model.MyCourses
	for _, inscripto := range inscriptos {
		courses = append(courses, inscripto.CourseId)
	}
	return courses
}
func (c *InscriptosClient) GetMyStudents(id uuid.UUID) model.StudentsInCourse {
	var inscriptos []model.Inscripto
	c.Db.Where("course_id = ?", id).Find(&inscriptos)
	var students model.StudentsInCourse
	for _, inscripto := range inscriptos {
		students = append(students, inscripto.UserId)
	}
	return students
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
	return count == 0, nil
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
