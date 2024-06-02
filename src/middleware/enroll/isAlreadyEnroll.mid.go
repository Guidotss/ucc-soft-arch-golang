package enroll

import (
	"fmt"
	"net/http"

	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func IsAlredyEnroll(service services.IInscriptionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("IsAlredyEnroll middleware")
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
			c.Abort()
			return
		}
		uid := userID.(uuid.UUID)
		fmt.Println("user_id: ", uid)

		courseID, exists := c.Get("courseID")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Course ID not found"})
			c.Abort()
			return
		}

		cid, err := uuid.Parse(courseID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Course ID"})
			c.Abort()
			return
		}
		existEnrrol, err := service.IsUserEnrolled(uid, cid)
		if existEnrrol {
			c.JSON(400, gin.H{"error": "User is already enrolled"})
			c.Abort()
			return
		}
		fmt.Println("Paso alredyEnrolled middleware")
		c.Next()
	}
}
