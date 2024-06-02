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

type StudentsInCourse []Student
