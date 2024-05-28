package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/courses"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/courses"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
)

type ICourseService interface {
	CreateCourse(courseDto dto.CreateCoursesRequestDto) dto.CreateCoursesResponseDto
	FindAllCourses() dto.GetAllCourses
	FindOneCourse(id uuid.UUID) dto.GetCourseDto
	UpdateCourse(dto dto.UpdateRequestDto) dto.UpdateResponseDto
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
		CourseImage:       courseDto.CourseImage,
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
		var courseDto dto.GetCourseDto
		courseDto.Id = result.Id
		courseDto.CategoryID = result.CategoryID
		courseDto.CourseName = result.CourseName
		courseDto.CourseDescription = result.CourseDescription
		courseDto.CoursePrice = result.CoursePrice
		courseDto.CourseDuration = result.CourseDuration
		courseDto.CourseCapacity = result.CourseCapacity
		courseDto.CourseInitDate = result.CourseInitDate
		courseDto.CourseState = result.CourseState
		courseDto.CourseImage = result.CourseImage

		allCoursesDto = append(allCoursesDto, courseDto)
	}
	return allCoursesDto
}

func (c *courseService) FindOneCourse(id uuid.UUID) dto.GetCourseDto {
	var result model.Course = c.client.GetById(id)
	var courseDto dto.GetCourseDto
	/*
	********FALTA MANEJO DE ERRORES*******
	 */
	courseDto.Id = result.Id
	courseDto.CategoryID = result.CategoryID
	courseDto.CourseName = result.CourseName
	courseDto.CourseDescription = result.CourseDescription
	courseDto.CoursePrice = result.CoursePrice
	courseDto.CourseDuration = result.CourseDuration
	courseDto.CourseCapacity = result.CourseCapacity
	courseDto.CourseInitDate = result.CourseInitDate
	courseDto.CourseState = result.CourseState
	courseDto.CourseImage = result.CourseImage

	return courseDto
}

func (c *courseService) UpdateCourse(newData dto.UpdateRequestDto) dto.UpdateResponseDto {
	var responseDto dto.UpdateResponseDto
	result := c.client.UpdateCourse(newData)
	responseDto.Id = result.Id
	responseDto.CategoryID = result.CategoryID
	responseDto.CourseName = result.CourseName
	responseDto.CourseDescription = result.CourseDescription
	responseDto.CoursePrice = result.CoursePrice
	responseDto.CourseDuration = result.CourseDuration
	responseDto.CourseCapacity = result.CourseCapacity
	responseDto.CourseInitDate = result.CourseInitDate
	responseDto.CourseState = result.CourseState
	responseDto.CourseImage = result.CourseImage

	return responseDto
}
