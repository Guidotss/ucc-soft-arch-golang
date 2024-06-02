package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/inscriptos"
	coursesDto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/courses"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/inscription"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
)

type IInscriptionService interface {
	Enroll(dto.EnrollRequestResponseDto) (dto.EnrollRequestResponseDto, error)
	GetMyCourses(uuid.UUID) (coursesDto.GetAllCourses, error)
	GetMyStudents(uuid.UUID) (dto.StudentsInCourse, error)
	IsUserEnrolled(userID uuid.UUID, courseID uuid.UUID) (bool, error)
	CourseExist(course_id uuid.UUID) (bool, error)
}

type inscriptionService struct {
	client inscriptos.InscriptosClient
}

func NewInscriptionService(client *inscriptos.InscriptosClient) IInscriptionService {
	return &inscriptionService{client: *client}
}

func (c *inscriptionService) Enroll(data dto.EnrollRequestResponseDto) (dto.EnrollRequestResponseDto, error) {
	var newEnroll = model.Inscripto{
		CourseId: data.CourseId,
		UserId:   data.UserId,
	}
	enroll, err := c.client.Enroll(newEnroll)
	if err != nil {
		return dto.EnrollRequestResponseDto{}, err
	}
	return dto.EnrollRequestResponseDto{
		CourseId: enroll.CourseId,
		UserId:   enroll.UserId,
	}, nil
}

func (c *inscriptionService) GetMyCourses(id uuid.UUID) (coursesDto.GetAllCourses, error) {
	response, err := c.client.GetMyCourses(id)
	if err != nil {
		return nil, err
	}
	var courses coursesDto.GetAllCourses
	for _, data := range response {
		course := coursesDto.GetCourseDto{
			Id:          data.Id,
			CourseName:  data.CourseName,
			CourseImage: data.CourseImage,
		}
		courses = append(courses, course)
	}
	return courses, nil
}
func (c *inscriptionService) GetMyStudents(id uuid.UUID) (dto.StudentsInCourse, error) {
	response, err := c.client.GetMyStudents(id)
	if err != nil {
		return nil, err
	}
	var students dto.StudentsInCourse
	for _, data := range response {
		studentDto := dto.Student{
			UserId:   data.Id,
			UserName: data.Name,
			Avatar:   data.Avatar,
		}
		students = append(students, studentDto)
	}
	return students, nil
}

func (c *inscriptionService) IsUserEnrolled(userID uuid.UUID, courseID uuid.UUID) (bool, error) {
	return c.client.IsUserEnrolled(userID, courseID)
}
func (c *inscriptionService) CourseExist(course_id uuid.UUID) (bool, error) {
	return c.client.CourseExist(course_id)
}
