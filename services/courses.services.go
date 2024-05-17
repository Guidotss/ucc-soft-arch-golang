package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/clients/courses"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/domain/dtos/courses"
	"github.com/Guidotss/ucc-soft-arch-golang.git/model"
)

type courseService struct{}

type ICourseService interface {
	CreateCourse(courseDto dto.CreateCoursesRequestDto) dto.CreateCoursesResponseDto
}

func NewCourseService() ICourseService {
	return &courseService{}
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

	var createdCourse = courses.Create(newCourse)

	return dto.CreateCoursesResponseDto{
		CourseName: createdCourse.CourseName,
		CourseId:   int(createdCourse.ID),
	}
}

/*
func Create(c courses.CreateCoursesRequestDto) *courses.CreateCoursesResponseDto {

}
*/
