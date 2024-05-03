package courses

type CoursesCreateRequest struct {
	CourseName        string  `json:"course_name"`
	CourseDescription string  `json:"course_description"`
	CoursePrice       float64 `json:"course_price"`
	CourseDuration    int     `json:"course_duration"`
	CourseImage       string  `json:"course_image"`
}

type CoursesCreateResponse struct {
	CourseName string `json:"course_name"`
	Message    string `json:"message"`
	CourseId   int    `json:"course_id"`
}
