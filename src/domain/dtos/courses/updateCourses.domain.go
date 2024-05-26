package courses

import "github.com/google/uuid"

type UpdateRequestDto struct {
	Id                uuid.UUID  `json:"id"`
	CourseName        *string    `json:"course_name"`
	CourseDescription *string    `json:"description"`
	CoursePrice       *float64   `json:"price"`
	CourseDuration    *int       `json:"duration"`
	CourseCapacity    *int       `json:"capacity"`
	CategoryID        *uuid.UUID `json:"category_id"`
	CourseInitDate    *string    `json:"init_date"`
	CourseState       *bool      `json:"state"`
	CourseImage       *string    `json:"image"`
}
type UpdateResponseDto struct {
	Id                uuid.UUID `json:"id"`
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
