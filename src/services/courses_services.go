package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/courses"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/courses"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type ICourseService interface {
	CreateCourse(courseDto dto.CreateCoursesRequestDto) dto.CreateCoursesResponseDto
	FindAllCourses() dto.GetAllCourses
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

func (c *courseService) FindAllCourses() dto.GetAllCourses {
	var courses model.Courses = c.client.GetAll()
	var allCoursesDto dto.GetAllCourses
	for _, result := range courses {
		var courseDto dto.GetAllCoursesResponseDto
		courseDto.Id = result.Id
		courseDto.CategoryID = result.CategoryID
		courseDto.CourseName = result.CourseName
		courseDto.CourseDescription = result.CourseDescription
		courseDto.CoursePrice = result.CoursePrice
		courseDto.CourseDuration = result.CourseDuration
		courseDto.CourseCapacity = result.CourseCapacity
		courseDto.CourseInitDate = result.CourseInitDate
		courseDto.CourseState = result.CourseState
		allCoursesDto = append(allCoursesDto, courseDto)
	}
	return allCoursesDto
}
