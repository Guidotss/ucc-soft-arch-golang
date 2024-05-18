package courses

type CreateCoursesRequestDto struct {
	CourseName        string  `json:"course_name"`
	CourseDescription string  `json:"course_description"`
	CoursePrice       float64 `json:"course_price"`
	CourseDuration    int     `json:"course_duration"`
	CourseCapacity    int     `json:"course_capacity"`
	CategoryID        string  `json:"category_id"`
	CourseInitDate    string  `json:"init_date"`
	CourseState       bool    `json:"state"`
}

type CreateCoursesResponseDto struct {
	CourseName string `json:"course_name"`
	CourseId   int    `json:"course_id"`
}
