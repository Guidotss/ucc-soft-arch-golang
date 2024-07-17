package course

import "github.com/gin-gonic/gin"

func CheckCourseId() gin.HandlerFunc {
	return func(c *gin.Context) {

		courseId := c.Param("id")
		if courseId == "" {
			c.JSON(400, gin.H{"error": "Course ID is required"})
			c.Abort()
			return
		}
		c.Set("courseId", courseId)
		c.Next()
	}
}
