package enroll

import (
	"fmt"

	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/inscription"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CourseExist(service services.IInscriptionService) gin.HandlerFunc {
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

		exist, err := service.CourseExist(course_id)
		if !exist {
			c.JSON(400, gin.H{"error": "Course doesn't exist"})
			c.Abort()
			return
		}
		if err != nil {
			c.Error(err)
		}
		c.Set("courseID", courseRequestString.CourseId)
		fmt.Println("Paso el Exist middleware")
		c.Next()
	}
}
