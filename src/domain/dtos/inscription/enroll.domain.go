package inscription

import "github.com/google/uuid"

type EnrollRequestResponseDto struct {
	CourseId uuid.UUID `json:"course_id"`
	UserId   uuid.UUID `json:"user_id"`
}
type Student struct {
	UserId uuid.UUID `json:"user_id"`
}
type Course struct {
	CourseId uuid.UUID `json:"course_id"`
}

type MyCoursesDto []Course
type StudentsInCourse []Student
