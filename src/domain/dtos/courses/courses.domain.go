package courses

type CreateCoursesRequestDto struct {
	CourseName        string  `json:"course_name"`
	CourseDescription string  `json:"description"`
	CoursePrice       float64 `json:"price"`
	CourseDuration    int     `json:"duration"`
	CourseCapacity    int     `json:"capacity"`
	CategoryID        string  `json:"category_id"`
	CourseInitDate    string  `json:"init_date"`
	CourseState       bool    `json:"state"`
}

type CreateCoursesResponseDto struct {
	CourseName string `json:"course_name"`
	CourseId   int    `json:"course_id"`
}
