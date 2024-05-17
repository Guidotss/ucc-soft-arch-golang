package services

import "github.com/Guidotss/ucc-soft-arch-golang.git/domain/dtos/courses"

func Create(c courses.CoursesCreateRequest) *courses.CoursesCreateResponse {
	return &courses.CoursesCreateResponse{
		Message:    "Create",
		CourseName: c.CourseName,
		CourseId:   12345,
	}
}
