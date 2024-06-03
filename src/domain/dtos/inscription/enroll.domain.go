package inscription

import "github.com/google/uuid"

type EnrollRequestResponseDto struct {
	CourseId uuid.UUID `json:"course_id"`
	UserId   uuid.UUID `json:"user_id"`
}
type Student struct {
	UserId   uuid.UUID `json:"user_id"`
	UserName string    `json:"user_name"`
	Avatar   string    `json:"avatar"`
}

type CourseIdString struct {
	CourseId string `json:"course_id"`
}
type MyCourse struct {
	Id          uuid.UUID `json:"course_id"`
	CourseName  string    `json:"course_name"`
	CourseImage string    `json:"course_image"`
}

type StudentsInCourse []Student
type MyCourses []MyCourse
