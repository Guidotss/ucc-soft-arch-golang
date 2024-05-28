package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/inscriptos"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/inscription"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
)

type IInscriptionService interface {
	Enroll(dto.EnrollRequestResponseDto) dto.EnrollRequestResponseDto
	GetMyCourses(uuid.UUID) dto.MyCoursesDto
	GetMyStudents(uuid.UUID) dto.StudentsInCourse
	IsUserEnrolled(userID uuid.UUID, courseID uuid.UUID) (bool, error)
	CourseExist(course_id uuid.UUID) (bool, error)
}

type inscriptionService struct {
	client inscriptos.InscriptosClient
}

func NewInscriptionService(client *inscriptos.InscriptosClient) IInscriptionService {
	return &inscriptionService{client: *client}
}

func (c *inscriptionService) Enroll(data dto.EnrollRequestResponseDto) dto.EnrollRequestResponseDto {
	var newEnroll = model.Inscripto{
		CourseId: data.CourseId,
		UserId:   data.UserId,
	}
	enroll := c.client.Enroll(newEnroll)

	return dto.EnrollRequestResponseDto{
		CourseId: enroll.CourseId,
		UserId:   enroll.UserId,
	}
}

func (c *inscriptionService) GetMyCourses(id uuid.UUID) dto.MyCoursesDto {
	courses := c.client.GetMyCourses(id)
	var coursesDto dto.MyCoursesDto
	for _, course := range courses {
		var courseDto dto.Course
		courseDto.CourseId = course
		coursesDto = append(coursesDto, courseDto)
	}
	return coursesDto
}
func (c *inscriptionService) GetMyStudents(id uuid.UUID) dto.StudentsInCourse {
	students := c.client.GetMyStudents(id)
	var studentsDto dto.StudentsInCourse
	for _, student := range students {
		var studentDto dto.Student
		studentDto.UserId = student
		studentsDto = append(studentsDto, studentDto)
	}
	return studentsDto
}

func (c *inscriptionService) IsUserEnrolled(userID uuid.UUID, courseID uuid.UUID) (bool, error) {
	return c.client.IsUserEnrolled(userID, courseID)
}
func (c *inscriptionService) CourseExist(course_id uuid.UUID) (bool, error) {
	return c.client.CourseExist(course_id)
}
