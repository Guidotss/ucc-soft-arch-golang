package services

import (
	"fmt"

	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/courses"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/courses"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
)

type ICourseService interface {
	CreateCourse(courseDto dto.CreateCoursesRequestDto) (dto.CreateCoursesResponseDto, error)
	FindAllCourses(filter string) (dto.GetAllCourses, error)
	FindOneCourse(id uuid.UUID) (dto.GetCourseDto, error)
	UpdateCourse(dto dto.UpdateRequestDto) (dto.UpdateResponseDto, error)
	DeleteCourse(id uuid.UUID) error
}

type courseService struct {
	client courses.CourseClient
}

func NewCourseService(client *courses.CourseClient) ICourseService {
	return &courseService{client: *client}
}

func (c *courseService) CreateCourse(courseDto dto.CreateCoursesRequestDto) (dto.CreateCoursesResponseDto, error) {

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

	createdCourse, err := c.client.Create(newCourse)
	if err != nil {
		return dto.CreateCoursesResponseDto{}, err
	}

	return dto.CreateCoursesResponseDto{
		CourseName: createdCourse.CourseName,
		CourseId:   createdCourse.Id,
	}, nil
}

func (c *courseService) FindAllCourses(filter string) (dto.GetAllCourses, error) {
	courses, err := c.client.GetAll(filter)
	if err != nil {
		return nil, err
	}
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
		courseDto.CourseCategoryName = result.Category.CategoryName
		courseDto.RatingAvg = result.RatingAvg
		allCoursesDto = append(allCoursesDto, courseDto)
	}
	return allCoursesDto, nil
}

func (c *courseService) FindOneCourse(id uuid.UUID) (dto.GetCourseDto, error) {
	result, err := c.client.GetById(id)
	if err != nil {
		return dto.GetCourseDto{}, err
	}
	return dto.GetCourseDto{
		Id:                 result.Id,
		CategoryID:         result.CategoryID,
		CourseName:         result.CourseName,
		CourseDescription:  result.CourseDescription,
		CoursePrice:        result.CoursePrice,
		CourseDuration:     result.CourseDuration,
		CourseCapacity:     result.CourseCapacity,
		CourseInitDate:     result.CourseInitDate,
		CourseState:        result.CourseState,
		CourseImage:        result.CourseImage,
		CourseCategoryName: result.Category.CategoryName,
		RatingAvg:          result.RatingAvg,
	}, nil
}

func (c *courseService) UpdateCourse(newData dto.UpdateRequestDto) (dto.UpdateResponseDto, error) {
	var course model.Course
	if newData.CourseName != nil {
		course.CourseName = *newData.CourseName
	}
	if newData.CourseDescription != nil {
		course.CourseDescription = *newData.CourseDescription
	}
	if newData.CoursePrice != nil {
		course.CoursePrice = *newData.CoursePrice
	}
	if newData.CourseDuration != nil {
		course.CourseDuration = *newData.CourseDuration
	}
	if newData.CourseCapacity != nil {
		course.CourseCapacity = *newData.CourseCapacity
	}
	if newData.CategoryID != nil {
		course.CategoryID = *newData.CategoryID
	}
	if newData.CourseInitDate != nil {
		course.CourseInitDate = *newData.CourseInitDate
	}
	if newData.CourseState != nil {
		course.CourseState = *newData.CourseState
	}
	if newData.CourseImage != nil {
		course.CourseImage = *newData.CourseImage
	}
	course.Id = newData.Id

	fmt.Println("UpdateCourse Service: ", course)

	result, err := c.client.UpdateCourse(course)
	if err != nil {
		return dto.UpdateResponseDto{}, err
	}
	return dto.UpdateResponseDto{
		Id:                result.Id,
		CategoryID:        result.CategoryID,
		CourseName:        result.CourseName,
		CourseDescription: result.CourseDescription,
		CoursePrice:       result.CoursePrice,
		CourseDuration:    result.CourseDuration,
		CourseCapacity:    result.CourseCapacity,
		CourseInitDate:    result.CourseInitDate,
		CourseState:       result.CourseState,
		CourseImage:       result.CourseImage,
	}, nil
}

func (c *courseService) DeleteCourse(id uuid.UUID) error {
	err := c.client.DeleteCourse(id)
	if err != nil {
		return err
	}
	return nil
}
