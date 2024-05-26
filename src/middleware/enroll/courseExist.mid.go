package enroll

import (
	"fmt"

	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/inscriptions"
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/inscription"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CourseExist(controller controller.InscriptionController) gin.HandlerFunc {
	return func(c *gin.Context) {
		var courseRequestString dto.CourseIdString
		err := c.BindJSON(&courseRequestString)
		if err != nil {
			fmt.Println("Error al pasar el UUID")
			c.JSON(400, gin.H{"error": "Invalid request format"})
			c.Abort()
			return
		}

		course_id, err := uuid.Parse(courseRequestString.CourseId)
		if err != nil {
			fmt.Println("Error al parsear el UUID")
			c.JSON(400, gin.H{"error": "Invalid UUID format"})
			c.Abort()
			return
		}

		if !controller.CourseExist(course_id) {
			c.JSON(400, gin.H{"error": "Course doesn't exist"})
			c.Abort()
			return
		}
		c.Set("courseID", courseRequestString.CourseId)
		fmt.Println("Paso el Exist middleware")
		c.Next()
	}
}
