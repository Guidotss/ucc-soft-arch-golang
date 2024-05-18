package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/courses"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/courses"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type ICourseService interface {
	CreateCourse(courseDto dto.CreateCoursesRequestDto) dto.CreateCoursesResponseDto
}

type courseService struct {
	client courses.CourseClient
}

func NewCourseService(client *courses.CourseClient) ICourseService {
	return &courseService{client: *client}
}

func (c *courseService) CreateCourse(courseDto dto.CreateCoursesRequestDto) dto.CreateCoursesResponseDto {

	var newCourse = model.Course{
		CourseName:        courseDto.CourseName,
		CourseDescription: courseDto.CourseDescription,
		CoursePrice:       courseDto.CoursePrice,
		CourseDuration:    courseDto.CourseDuration,
		CourseCapacity:    courseDto.CourseCapacity,
		CategoryID:        courseDto.CategoryID,
		CourseInitDate:    courseDto.CourseInitDate,
		CourseState:       courseDto.CourseState,
	}

	createdCourse := c.client.Create(newCourse)

	return dto.CreateCoursesResponseDto{
		CourseName: createdCourse.CourseName,
		CourseId:   createdCourse.Id,
	}
}
