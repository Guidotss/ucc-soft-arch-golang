package courses

import "github.com/google/uuid"

type CreateCoursesRequestDto struct {
	CourseName        string    `json:"course_name"`
	CourseDescription string    `json:"description"`
	CoursePrice       float64   `json:"price"`
	CourseDuration    int       `json:"duration"`
	CourseCapacity    int       `json:"capacity"`
	CategoryID        uuid.UUID `json:"category_id"`
	CourseInitDate    string    `json:"init_date"`
	CourseState       bool      `json:"state"`
	CourseImage       string    `json:"image"`
}

type CreateCoursesResponseDto struct {
	CourseName string    `json:"course_name"`
	CourseId   uuid.UUID `json:"course_id"`
}

type GetOneCourseRequestDto struct {
	CourseId uuid.UUID `json:"course_id"`
}

type GetOneCourseResponseDto struct {
	CategoryID        uuid.UUID `json:"category_id"`
	CourseName        string    `json:"course_name"`
	CourseDescription string    `json:"description"`
	CoursePrice       float64   `json:"price"`
	CourseDuration    int       `json:"duration"`
	CourseCapacity    int       `json:"capacity"`
	CourseInitDate    string    `json:"init_date"`
	CourseState       bool      `json:"state"`
	CourseImage       string    `json:"image"`
}
