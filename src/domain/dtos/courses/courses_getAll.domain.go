package courses

import "github.com/google/uuid"

type GetAllCoursesResponseDto struct {
	Id                uuid.UUID `json:"id"`
	CategoryID        uuid.UUID `json:"category_id"`
	CourseName        string    `json:"course_name"`
	CourseDescription string    `json:"description"`
	CoursePrice       float64   `json:"price"`
	CourseDuration    int       `json:"duration"`
	CourseCapacity    int       `json:"capacity"`
	CourseInitDate    string    `json:"init_date"`
	CourseState       bool      `json:"state"`
}

type GetAllCourses []GetAllCoursesResponseDto
